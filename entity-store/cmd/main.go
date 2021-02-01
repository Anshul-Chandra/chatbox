package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	v1 "github.com/chatbox/entity-store/pkg/api/v1/user"
	"github.com/chatbox/proto/gen/v1/user"
)

func startGRPC() {
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// register user service
	user.RegisterUserServiceServer(grpcServer, &v1.UserService{})

	// TODO: register group service

	log.Println("gRPC server ready...")
	grpcServer.Serve(lis)
}

func main() {
	startGRPC()
}
