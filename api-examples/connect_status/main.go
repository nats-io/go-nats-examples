package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_status]
	nc, err := nats.Connect("nats://demo.nats.io:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	getStatusTxt := func(nc *nats.Conn) string {
		switch nc.Status() {
		case nats.CONNECTED:
			return "Connected"
		case nats.CLOSED:
			return "Closed"
		default:
			return "Other"
		}
	}
	log.Printf("The connection is %v\n", getStatusTxt(nc))

	nc.Close()

	log.Printf("The connection is %v\n", getStatusTxt(nc))

	// [end connect_status]
}
