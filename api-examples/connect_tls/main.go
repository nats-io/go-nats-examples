package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_tls]
	nc, err := nats.Connect("localhost",
		nats.ClientCert("resources/certs/cert.pem", "resources/certs/key.pem"),
		nats.RootCAs("resources/certs/ca.pem"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_tls]
}
