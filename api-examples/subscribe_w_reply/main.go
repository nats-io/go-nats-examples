package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin subscribe_w_reply]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe
	sub, err := nc.SubscribeSync("time")
	if err != nil {
		log.Fatal(err)
	}

	// Read a message
	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Get the time
	timeAsBytes := []byte(time.Now().String())

	// Send the time as the response.
	msg.Respond(timeAsBytes)

	// [end subscribe_w_reply]
}
