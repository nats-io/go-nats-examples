package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin publish_bytes]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	if err := nc.Publish("updates", []byte("All is Well")); err != nil {
		log.Fatal(err)
	}
	// Make sure the message goes through before we close
	nc.Flush()
	// [end publish_bytes]
}
