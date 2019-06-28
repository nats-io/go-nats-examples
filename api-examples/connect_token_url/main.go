package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_token_url]
	// Token in URL
	nc, err := nats.Connect("mytoken@localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_token_url]
}
