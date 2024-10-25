package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Can't connect to NATS, %v", err)
	}
	defer nc.Close()

	sub, err := nc.Subscribe("intros", processData) // callback function that handles the reception of a message
	if err != nil {
		log.Fatalf("Can't subscribe to NATS queue 'queueName', %v", err)
	}
	defer sub.Unsubscribe()

	// the subscription is non-blocking, that means that we need to find a wait for the server to keep up
	// if could be a for loop
	// In this example just a 1-hour wait
	time.Sleep(1 * time.Hour)

}

func processData(msg *nats.Msg) {
	log.Println("Got message", string(msg.Data))
}
