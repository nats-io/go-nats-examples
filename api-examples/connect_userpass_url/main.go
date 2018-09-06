package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_userpass_url]
	// Set a user and plain text password
	nc, err := nats.Connect("myname:password@localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_userpass_url]
}
