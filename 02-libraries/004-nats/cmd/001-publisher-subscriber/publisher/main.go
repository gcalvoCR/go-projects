package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gcalvoCR/go-learn/model"

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
		message := fmt.Sprintf(" %v", count)
		if err := nc.Publish("intros", []byte(data)); err != nil {
			log.Println("Error publishing the message")
		} else {
			log.Println("Message sent -->", message)
		}
		time.Sleep(2 * time.Second)
		count++
	}
}
