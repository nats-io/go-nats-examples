package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin publish_json]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Define the object
	type stock struct {
		Symbol string
		Price  int
	}

	// Publish the message
	if err := ec.Publish("updates", &stock{Symbol: "GOOG", Price: 1200}); err != nil {
		log.Fatal(err)
	}
	// [end publish_json]
}
