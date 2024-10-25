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

	js.Subscribe("orders.us", processMsg, nats.BindStream("ORDERS"))

	time.Sleep(10 * time.Second)

	log.Println("shutting down application...")
}

func processMsg(msg *nats.Msg) {
	fmt.Printf("INFO - Got reply - %s\n", string(msg.Data))
}
