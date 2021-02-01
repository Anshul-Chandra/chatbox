all: build

clean:
	@rm -rf bin

vendor:
	@go mod vendor
	@go mod tidy

build: clean vendor
	@go build -o bin/user-service -mod=vendor cmd/main.go

run:
	@go run cmd/main.go