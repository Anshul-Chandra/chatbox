module github.com/chatbox/entity-store

go 1.13

replace github.com/chatbox/proto => ../proto

require (
	github.com/chatbox/proto v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
)
