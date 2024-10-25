package main

import (
	"fmt"
	"os"
	"os/signal"

	"flag"

	"github.com/charmbracelet/log"
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
	log.Info("Connecting to", url)
	nc, err := nats.Connect(url)
	fatalOnErr(err)
	defer nc.Close()

	// 2 --> connect to jetstream connection
	js, err := nc.JetStream()
	fatalOnErr(err)

	// to list all the buckets and their keys
	for bucketName := range js.KeyValueStoreNames() {
		fmt.Printf("Bucket - %s\n", bucketName)

		kvBucket, err := js.KeyValue(bucketName)
		if err != nil {
			log.Printf("failed to bind to bucket: %s", err)
			continue
		}

		keys, err := kvBucket.Keys()
		if err != nil {
			continue
		}

		// listing the keys of a bucket
		for _, key := range keys {
			value, _ := kvBucket.Get(key)
			fmt.Printf("\t%s\t%s\n", key, value.Value())
		}
	}

	// create a new bucket call 'sensors'
	bucketName := "sensors"
	sensors, err := js.CreateKeyValue(&nats.KeyValueConfig{
		Bucket: bucketName,
	})
	fatalOnErr(err)

	// add a few key & value pairs
	sensors.PutString("temperature", "48deg")
	sensors.PutString("humidity", "50%")
	sensors.PutString("pressure", "10bars")

	monitor, err := sensors.Watch("*")
	fatalOnErr(err)

	go handleReadings(monitor.Updates())

	// cleanly exit application if signal is caught
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sensors.ListKeys()
}

func handleReadings(ch <-chan nats.KeyValueEntry) {
	for entry := range ch {
		if entry != nil {
			log.Info("new reading", "sensor", entry.Key(), "value", string(entry.Value()))
		}
	}
}
