module github.com/chatbox/grpc-gateway

go 1.13

require (
	github.com/chatbox/proto v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	google.golang.org/grpc v1.35.0
)

replace github.com/chatbox/proto => ../proto
