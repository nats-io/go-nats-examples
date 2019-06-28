package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_options]
	nc, err := nats.Connect("demo.nats.io", nats.Name("API Options Example"), nats.Timeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_options]
}
