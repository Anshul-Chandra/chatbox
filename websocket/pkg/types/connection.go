package types

import (
	"log"
	"time"

	"github.com/chatbox/websocket/pkg/uuid"
	ws "github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Connection is a wrapper arouund the websocket connection object
type Connection struct {
	connectionID string
	conn         *ws.Conn
	Send         chan *Message
	receive      chan *Message
}

// NewConnection returns a new instance of connection
func NewConnection(conn *ws.Conn) (*Connection, error) {
	connectionID, err := uuid.GetUUID()
	if err != nil {
		return nil, err
	}
	return &Connection{
		connectionID: connectionID,
		conn:         conn,
		Send:         make(chan *Message),
		receive:      make(chan *Message),
	}, nil
}

// GetID returns the ID associated with a connection
func (c *Connection) GetID() string {
	if c == nil {
		return ""
	}

	return c.connectionID
}

// GetWebsocketConnection returns the websocket connection that this object holds
func (c *Connection) GetWebsocketConnection() *ws.Conn {
	if c == nil {
		return nil
	}

	return c.conn
}

// Read reads a message from the websocket connection and writes to the write channel maintained by the server
func (c *Connection) Read(writeChannel chan *Message, unregister chan string) {
	if c == nil {
		log.Printf("connection object is nil")
		return
	}
	defer func() {
		log.Printf("unregistering client")
		unregister <- c.connectionID
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	log.Printf("started listening for messages for connection id %s", c.connectionID)
	for {
		message := &Message{}
		err := c.conn.ReadJSON(message)
		log.Printf("received message: %+v", message)
		if err != nil {
			log.Printf("error received: %s", err.Error())
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Printf("message received: %+v", message)

		writeChannel <- message
	}
}

func (c *Connection) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			log.Printf("Writing a new message for client %s. Message: %+v", c.connectionID, message)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(ws.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteJSON(message)
			if err != nil {
				log.Printf("Unable to write message. Error: %s", err.Error())
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(ws.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
