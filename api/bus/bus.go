package bus

import (
	"crypto/tls"
	"log"

	"github.com/nats-io/nats.go"
)

var conn *nats.Conn
var jsonConn *nats.EncodedConn

// Init the connection to the NATS server
func Init(url string, tlsConfig *tls.Config) error {
	log.Println("bus: init")
	var err error

	if tlsConfig != nil {
		conn, err = nats.Connect(url, nats.Secure(tlsConfig))
	} else {
		conn, err = nats.Connect(url)
	}
	if err != nil {
		return err
	}

	jsonConn, err = nats.NewEncodedConn(conn, nats.JSON_ENCODER)

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
