package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin connect_tls]
	nc, err := nats.Connect("nats://localhost:4222",
		nats.ClientCert("resources/certs/cert.pem", "resources/certs/key.pem"),
		nats.RootCAs("resources/certs/ca.pem"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_tls]
}
