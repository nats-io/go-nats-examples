package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin ping_20s]
	// Set Ping Interval to 20 seconds
	nc, err := nats.Connect("nats://demo.nats.io:4222", nats.PingInterval(20*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end ping_20s]
}
