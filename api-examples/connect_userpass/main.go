package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_userpass]
	// Set a user and plain text password
	nc, err := nats.Connect("127.0.0.1", nats.UserInfo("myname", "password"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_userpass]
}
