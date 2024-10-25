package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gcalvoCR/go-learn/model"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Can't connect to NATS, %v", err)
	}
	defer nc.Close()

	count := 0
	pl := &model.Payload{
		Data: "Hello world",
	}
	for {
		pl.Count = count
		// Encode payload
		data, _ := json.Marshal(pl)
		reply, err := nc.Request("intro", data, (500 * time.Millisecond)) // request needs a timeout
		time.Sleep(2 * time.Second)
		if err != nil {
			fmt.Println("Error sending message", err)
			continue
		}
		fmt.Printf("Sent %v, Got reply %v\n", pl.Count, string(reply.Data))
		count++
	}
}
