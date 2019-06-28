package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin publish_bytes]
	nc, err := nats.Connect("demo.nats.io", nats.Name("API PublishBytes Example"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	if err := nc.Publish("updates", []byte("All is Well")); err != nil {
		log.Fatal(err)
	}
	// [end publish_bytes]
}
