package bus

import (
	"log"

	"github.com/nats-io/nats.go"
)

var conn *nats.Conn

// Init the connection to the NATS server
func Init(url string) error {
	log.Println("bus: init")
	var err error
	conn, err = nats.Connect(url)
	return err
}

// Close the connection to the NATS server
func Close() {
	log.Println("bus: close")
	conn.Drain()
	conn.Close()
}

// Conn returns the connection to the NATS server
func Conn() *nats.Conn {
	return conn
}
