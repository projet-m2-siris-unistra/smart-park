package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	var err error

	natsURL, ok := os.LookupEnv("NATS_URL")
	if !ok {
		natsURL = nats.DefaultURL
	}

	nc, err = nats.Connect(natsURL)
	defer nc.Close()
	if err != nil {
		log.Fatalf("unable to connect to bus: %v", err)
	}

	http.HandleFunc("/", indexHandler)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	log.Printf("main: listening on port %s", port)
	http.ListenAndServe(":"+port, nil)
	log.Println("main: exiting")
}
