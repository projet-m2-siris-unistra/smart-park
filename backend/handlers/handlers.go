package handlers

import (
	"log"

	"github.com/nats-io/nats.go"
)

// Register the handlers
func Register(conn *nats.Conn) {
	log.Println("handlers: register")
	conn.Subscribe("ping", ping)
}

func Publisher(conn *nats.Conn, subj string, data []byte) {
	log.Println("handlers: publish")
	conn.Publish(subj, data)
}
