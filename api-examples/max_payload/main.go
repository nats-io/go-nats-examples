package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin max_payload]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	mp := nc.MaxPayload()
	log.Printf("Maximum payload is %v bytes", mp)

	// Do something with the max payload

	// [end max_payload]
}
