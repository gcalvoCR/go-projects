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

	sub, err := nc.Subscribe("intro", processData)
	if err != nil {
		log.Fatalf("Can't subscribe to NATS subject, %v", err)
	}
	defer sub.Unsubscribe()

	time.Sleep(1 * time.Minute)

}

func processData(msg *nats.Msg) {
	pl := &model.Payload{}
	json.Unmarshal(msg.Data, pl)
	replyData := fmt.Sprintf("Ack message # %v", pl.Count)
	msg.Respond([]byte(replyData))
	fmt.Println("Got message", string(msg.Data))
}
