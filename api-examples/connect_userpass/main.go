package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_userpass]
	// Set a user and plain text password
	nc, err := nats.Connect("nats://localhost:4222", nats.UserInfo("myname", "password"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_userpass]
}
