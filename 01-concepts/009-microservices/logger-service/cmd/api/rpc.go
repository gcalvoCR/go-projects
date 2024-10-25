package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// first declare a type that's going to be the RPC server
type RPCServer struct{}

// second we declare the type of data we are going to receive
type RPCPayload struct {
	Name string
	Data string
}

// declare the function
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Processed payload via RPC: " + payload.Name
	return nil
}
