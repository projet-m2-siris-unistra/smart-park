package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
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

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	zones.Register(api)
	tenants.Register(api)
	r.Use(utils.LoggingMiddleware)

	log.Fatal(http.ListenAndServe(":9123", r))
}
