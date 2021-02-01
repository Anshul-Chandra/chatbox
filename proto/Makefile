all: build

clean:
	@rm -rf gen/*
	@mkdir -p gen/v1/user
	@mkdir -p gen/v1/group
	@mkdir -p gen/v1/message
	@echo "\t==> clean complete!"

build: clean generate

generate: gen-user gen-group gen-message
	@echo "\t==> proto stub generation complete!"

gen-user:
	@echo "\t==> generating user service proto definitions. . ."
	@protoc -I/usr/local/include \
	-I$(GOPATH)/src/github.com \
	-I$(GOPATH)/src/github.com/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:gen/v1/user --grpc-gateway_out=logtostderr=true:gen/v1/user \
	--grpc-gateway_opt paths=source_relative \
	--go_opt=paths=source_relative \
	--proto_path v1/user user.proto user_service.proto

gen-group:
	@echo "\t==> generating group service proto definitions. . ."
	@protoc -I/usr/local/include \
	-I$(GOPATH)/src/github.com \
	-I$(GOPATH)/src/github.com/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:gen/v1/group --grpc-gateway_out=logtostderr=true:gen/v1/group \
	--grpc-gateway_opt paths=source_relative \
	--go_opt=paths=source_relative \
	--proto_path v1/group group.proto group_service.proto

gen-message:
	@echo "\t==> generating message proto definitions. . ."
	@protoc -I/usr/local/include \
	--go_out=plugins=grpc:gen/v1/message \
	--go_opt=paths=source_relative \
	--proto_path v1/message message.proto