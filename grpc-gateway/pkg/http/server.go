package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/chatbox/grpc-gateway/pkg/common"
	"github.com/chatbox/proto/gen/v1/user"
)

// Server is a http server
type Server struct {
	mux         *runtime.ServeMux
	httpPort    string
	connections map[string]grpc.ClientConnInterface
}

// NewServer returns a new instance of server
func NewServer(httpPort string) *Server {
	s := &Server{
		mux:      runtime.NewServeMux(),
		httpPort: httpPort,
	}
	s.initConnections()
	return s
}

// Run starts the http server
func (s *Server) Run() {
	if s == nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverAddress := fmt.Sprintf("localhost:%s", s.httpPort)
	log.Printf("starting http server at address: %s", serverAddress)

	// set up mux with reverse proxy
	s.RegisterClients(ctx)

	if err := http.ListenAndServe(serverAddress, s.mux); err != nil {
		log.Fatalf("Server stopped. Error: %+v", err)
	}
}

// GetConnections returns the grpc client connections configured in server
func (s *Server) GetConnections() map[string]grpc.ClientConnInterface {
	if s == nil {
		return nil
	}
	return s.connections
}

// GetMux returns the grpc server mux configured in server
func (s *Server) GetMux() *runtime.ServeMux {
	if s == nil {
		return nil
	}
	return s.mux
}

func (s *Server) initConnections() {
	if s == nil {
		return
	}
	s.connections = make(map[string]grpc.ClientConnInterface)
	conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	s.connections[common.EntityStore] = conn
	log.Printf("connections: %+v", s.connections)
}

// RegisterClients registers grpc gateways for different microservices
func (s *Server) RegisterClients(ctx context.Context) error {
	if s == nil {
		return nil
	}
	for _, conn := range s.GetConnections() {
		// Register user service client
		client := user.NewUserServiceClient(conn)
		if err := user.RegisterUserServiceHandlerClient(ctx, s.GetMux(), client); err != nil {
			return err
		}
	}

	return nil
}
