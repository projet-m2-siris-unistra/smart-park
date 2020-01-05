package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"

	"github.com/projet-m2-siris-unistra/smart-park/api/bus"
	"github.com/projet-m2-siris-unistra/smart-park/api/utils"
	"github.com/projet-m2-siris-unistra/smart-park/api/v1/tenants"
	"github.com/projet-m2-siris-unistra/smart-park/api/v1/zones"
)

func main() {
	natsURL, ok := os.LookupEnv("NATS_URL")
	if !ok {
		natsURL = nats.DefaultURL
	}

	err := bus.Init(natsURL)
	defer bus.Close()

	if err != nil {
		log.Fatalf("unable to connect to bus: %v", err)
	}

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	zones.Register(api)
	tenants.Register(api)
	r.Use(utils.LoggingMiddleware)

	log.Fatal(http.ListenAndServe(":9123", r))
}
