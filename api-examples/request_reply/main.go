package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin request_reply]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Send the request
	msg, err := nc.Request("time", nil, time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Use the response
	log.Printf("Reply: %s", msg.Data)

	// Close the connection
	nc.Close()
	// [end request_reply]
}
