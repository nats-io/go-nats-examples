package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin servers_added]
	// Be notified if a new server joins the cluster.
	// Print all the known servers and the only the ones that were discovered.
	nc, err := nats.Connect("demo.nats.io",
		nats.DiscoveredServersHandler(func(nc *nats.Conn) {
			log.Printf("Known servers: %v\n", nc.Servers())
			log.Printf("Discovered servers: %v\n", nc.DiscoveredServers())
		}))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Do something with the connection

	// [end servers_added]
}
