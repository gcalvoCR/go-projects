package main

import (
	"fmt"
	"log"
	"time"

	"flag"

	"github.com/nats-io/nats.go"
)

var (
	username string
	password string
	hostname = "localhost"
	port     = 4222
)

func init() {
	flag.StringVar(&username, "u", username, "username for NATS Server")
	flag.StringVar(&password, "p", password, "password for NATS Server")
	flag.StringVar(&hostname, "host", hostname, "NATS Server hostnamer")
	flag.IntVar(&port, "port", port, "NATS Server port")
	flag.Parse()
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := fmt.Sprintf("nats://%v:%v", hostname, port)
	if username != "" {
		url = fmt.Sprintf("nats://%v:%v@%v:%v", username, password, hostname, port)
	}

	// 1 --> connect to nats server
	log.Println("Connecting to", url)
	nc, err := nats.Connect(url)
	fatalOnErr(err)
	defer nc.Close()

	// 2 --> connect to jetstream connection
	js, err := nc.JetStream()
	fatalOnErr(err)

	// 3 --> we create a consumer
	_, err = js.AddConsumer("ORDERS", &nats.ConsumerConfig{
		Durable:      "my-consumer-1",                    // the name we want to give the consumer
		Description:  "This is an example of a consumer", // this shows up when you list your consumers
		ReplayPolicy: nats.ReplayInstantPolicy,
	})
	fatalOnErr(err)
	// Note:
	// Consumers are Idepondent,
	// running multiple instances of this code will now try to recreate another consumer

	// 4 --> we the subscribe using the consumer
	sub, err := js.PullSubscribe("orders.us", "my-consumer-1")
	fatalOnErr(err)
	go processMsg(sub)

	time.Sleep(10 * time.Second)
	sub.Unsubscribe()

	log.Println("shutting down application...")
}

func processMsg(sub *nats.Subscription) {
	for sub.IsValid() { // we are pulling messages from nats via this consumer
		// we could get as many messages as we want, in this case just 1
		msgs, err := sub.Fetch(1)
		if err == nil {
			fmt.Printf("INFO - Got reply - %s\n", string(msgs[0].Data))
		}
	}
}
