package main

import (
	server "github.com/chatbox/grpc-gateway/pkg/http"
)

const httpPort = "8080"

func main() {
	s := server.NewServer(httpPort)
	s.Run()
}
