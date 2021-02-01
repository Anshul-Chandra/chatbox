package websocket

import (
	"log"
	"net/http"

	"github.com/chatbox/websocket/pkg/store"
	"github.com/chatbox/websocket/pkg/types"
	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Server is the chat server that handles websocket traffic
type Server struct {
	clientStore store.Store

	// Inbound messages from the clients.
	broadcast chan *types.Message

	// Unregister requests from clients.
	unregister chan string
}

// NewServer returns a new instance of Server
func NewServer() *Server {
	return &Server{
		clientStore: store.NewInMemoryStore(),
		broadcast:   make(chan *types.Message),
		unregister:  make(chan string),
	}
}

// Run starts the websockets server
func (s *Server) Run() {
	for {
		select {
		case clientID := <-s.unregister:
			err := s.clientStore.Unregister(clientID)
			if err != nil {
				log.Printf("error occured while unregistering a client. Error: %s", err.Error())
			}
		case message := <-s.broadcast:
			if message != nil {
				target := message.ReceiverID
				clients, err := s.clientStore.Get(target)
				if err != nil {
					log.Printf("error occured while getting clients for id: %s. Error: %s", target, err.Error())
				}
				for _, client := range clients {
					client.Send <- message
				}
			}
		}
	}
}

// WebsocketHandler is the handler for new websocket requests
func (s *Server) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("registering a new connection")
	// Create a new client connection and register it
	clientConnection, err := types.NewConnection(conn)
	if err != nil {
		log.Printf("unable to generate a new connection object. Error: %s", err.Error())
		return
	}
	err = s.clientStore.Register(clientConnection)
	if err != nil {
		log.Printf("error occured while registering client. Error: %s", err.Error())
		return
	}

	log.Printf("successfully registered a new connection. Conection ID: %s", clientConnection.GetID())

	// spawn new goroutines for readers and writers
	go clientConnection.Write()
	go clientConnection.Read(s.broadcast, s.unregister)
}
