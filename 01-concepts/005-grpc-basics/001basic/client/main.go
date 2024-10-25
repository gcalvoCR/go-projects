package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gcalvoCR/go-learn/cmd/004-grpc/001basic/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:8080"
)

func main() {
	flag.StringVar(&address, "a", address, "gRPC server address host:port")
	flag.Parse()

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.NewClient(address, opts)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to connect to grpc service: %v", err))
	}
	defer conn.Close()

	client := model.NewMathServiceClient(conn)

	ctx := context.Background()
	in := &model.MathRequest{Operand1: 11, Operant2: 20}
	result, err := client.Add(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("add rpc failed: %v", err))
	}

	fmt.Printf("Add(%v) =>%v\n", in, result)

	result, err = client.Sub(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("sub rpc failed: %v", err))
	}

	fmt.Printf("Sub(%v) =>%v\n", in, result)
}
