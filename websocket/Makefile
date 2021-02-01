all: build

clean:
	@rm -rf bin
	@echo "\t==> clean complete!"

build: clean build-websocket-server

build-websocket-server:
	@echo "\t==> building websocket server binary. . ."
	@go build -o bin/websocket-server cmd/main.go
	@echo "\t==> build complete!"

run-server:
	go run cmd/main.go