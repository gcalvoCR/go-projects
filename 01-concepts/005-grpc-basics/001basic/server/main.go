package main

import (
	"flag"
	"log"
	"net"

	"github.com/gcalvoCR/go-learn/cmd/004-grpc/001basic/model"
	"google.golang.org/grpc"
)

var (
	address = "localhost:8080"
)

func main() {
	flag.StringVar(&address, "a", address, "gRPC server address host:port")
	flag.Parse()

	// create a socket to listen on
	ls, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to start grpc server on address %v:%v", address, err)
	}
	// create a new grpc server
	s := grpc.NewServer()
	// register the implementation that we have
	model.RegisterMathServiceServer(s, &server{})
	log.Printf("server up and running: %s", address)

	if err := s.Serve(ls); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
