package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq" //we will be utilizing to interact with our database
	"github.com/nats-io/nats.go"

	"github.com/projet-m2-siris-unistra/smart-park/backend/bus"
	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
	"github.com/projet-m2-siris-unistra/smart-park/backend/handlers"
)

func loadTLS(prefix string) (*tls.Config, error) {
	certFile, ok := os.LookupEnv(prefix + "_CERT")
	if !ok {
		return nil, nil
	}
	keyFile, ok := os.LookupEnv(prefix + "_KEY")
	if !ok {
		return nil, nil
	}
	caFile, ok := os.LookupEnv(prefix + "_CA")
	if !ok {
		return nil, nil
	}

	// Load client certificate
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Load CA
	pool := x509.NewCertPool()
	rootPEM, err := ioutil.ReadFile(caFile)
	if err != nil || rootPEM == nil {
		return nil, fmt.Errorf("error loading or parsing rootCA file: %v", err)
	}
	ok = pool.AppendCertsFromPEM(rootPEM)
	if !ok {
		return nil, fmt.Errorf("failed to parse root certificate from %q", caFile)
	}

	return &tls.Config{
		MinVersion:   tls.VersionTLS12,
		RootCAs:      pool,
		Certificates: []tls.Certificate{cert},
	}, nil
}

func main() {
	var err error

	// Connect to the DB - defini dans docker-compose
	databaseURL, ok := os.LookupEnv("DATABASE")
	if !ok {
		databaseURL = "postgres:///postgres?sslmode=disable"
	}

	err = database.Init(databaseURL)
	defer database.Close()

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	natsURL, ok := os.LookupEnv("NATS_URL")
	if !ok {
		natsURL = nats.DefaultURL
	}

	tlsConfig, err := loadTLS("NATS")
	if err != nil {
		log.Fatalf("unable to load certificates: %v", err)
	}

	err = bus.Init(natsURL, tlsConfig)
	defer bus.Close()

	if err != nil {
		log.Fatalf("unable to connect to bus: %v", err)
	}

	// Register the handlers
	handlers.Register(bus.Conn())

	// wait signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Receive from c
	log.Println("main: exiting")
}
