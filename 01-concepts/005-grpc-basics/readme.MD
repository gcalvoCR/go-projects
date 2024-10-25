# gRPC

You are describing a service that allows you to call that are exposed in that service remotely.

- We first describe our service using .proto
- We compile it (using gRPC compiler)
- You can implement the services and also call the services

1. Client Library
2. Server Stubs



### Notes:
in gRPC you can only have one message as input and one message as return

### dependencies
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

###  Docs
https://grpc.io/docs/languages/go/quickstart/

```sh
go get -u google.golang.org/grpc

# wrong command
protoc --go_out=. cmd/004-grpc/model/messages.proto

protoc --go-grpc_out=. cmd/004-grpc/model/messages.proto

```
