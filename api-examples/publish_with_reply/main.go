package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin publish_with_reply]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create a unique subject name
	uniqueReplyTo := nats.NewInbox()

	// Listen for a single response
	sub, err := nc.SubscribeSync(uniqueReplyTo)
	if err != nil {
		log.Fatal(err)
	}

	// Send the request
	if err := nc.PublishRequest("time", uniqueReplyTo, nil); err != nil {
		log.Fatal(err)
	}

	// Read the reply
	msg, err := sub.NextMsg(time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Use the response
	log.Printf("Reply: %s", msg.Data)

	// Close the connection
	nc.Close()
	// [end publish_with_reply]
}
