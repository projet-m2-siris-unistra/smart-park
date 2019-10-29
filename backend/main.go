package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"

	"github.com/projet-m2-siris-unistra/smart-park/backend/bus"
	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
	"github.com/projet-m2-siris-unistra/smart-park/backend/handlers"
)

func main() {
	var err error

	// Connect to the DB
	databaseURL, ok := os.LookupEnv("DATABASE")
	if !ok {
		databaseURL = "postgres:///postgres?sslmode=disable"
	}

	err = database.Init(databaseURL)
	defer database.Close()

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	// Connect to NATS
	natsURL, ok := os.LookupEnv("NATS_URL")
	if !ok {
		natsURL = nats.DefaultURL
	}

	err = bus.Init(natsURL)
	defer bus.Close()

	if err != nil {
		log.Fatalf("unable to connect to bus: %v", err)
	}

	// Register the handlers
	handlers.Register(bus.Conn())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("main: exiting")
}
