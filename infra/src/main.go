package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/go-nats"
)

// JSON : https://github.com/tidwall/gjson

type message struct {
	server  string
	subject string
	data    string
	timeout uint
}

var nc *nats.Conn
var m message

func main() {

	var err error

	m.subject = flag.Arg(0)
	m.data = "test"
	m.timeout = 1
	m.server = nats.DefaultURL

	mode := "pubsub"
	// Connect to NATS
	nc, err := nats.Connect(m.server, nats.PingInterval(20*time.Second))

	panic(err, m)
	getStatusTxt := func(nc *nats.Conn) string {
		switch nc.Status() {
		case nats.CONNECTED:
			return "Connected"
		case nats.CLOSED:
			return "Closed"
		default:
			return "Other"
		}

	}
	log.Printf("The connection is %v\n", getStatusTxt(nc))

	switch mode {
	case "pub", "publish":
		err = publish(m.subject, m.data)
	case "sub", "subscribe":
		err = subscribe(m.subject)
	case "pubsub":
		err = pubsub(m.subject, m.data)
	default:
		flag.Usage()
	}

	nc.Close()
	log.Printf("The connection is %v\n", getStatusTxt(nc))
	os.Exit(0)

}

func panic(err error, m message) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s - %s < err: %v\n", m.server, m.subject, err)
		os.Exit(1)
	}
	defer nc.Close()
}
