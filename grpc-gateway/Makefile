all: build

clean:
	@rm -rf bin

build: clean
	@go build -o bin/grpc-gateway -mod=vendor cmd/main.go

run:
	@go run cmd/main.go