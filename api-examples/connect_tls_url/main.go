package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin connect_tls_url]
	nc, err := nats.Connect("tls://localhost", nats.RootCAs("resources/certs/ca.pem")) // May need this if server is using self-signed certificate
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end connect_tls_url]
}
