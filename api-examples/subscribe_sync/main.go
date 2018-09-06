package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin subscribe_sync]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe
	sub, err := nc.SubscribeSync("updates")
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a message
	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Use the response
	log.Printf("Reply: %s", msg.Data)

	// Close the connection
	nc.Close()
	// [end subscribe_sync]
}
