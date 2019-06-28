package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_status]
	nc, err := nats.Connect("demo.nats.io", nats.Name("API Example"))
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
