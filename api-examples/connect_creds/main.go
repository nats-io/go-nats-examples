package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_creds]
	nc, err := nats.Connect("localhost", nats.UserCredentials("path_to_creds_file"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_creds]
}
