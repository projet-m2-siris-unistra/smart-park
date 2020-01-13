package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
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

type deviceInfos struct {
	DeviceName string                        `json:"deviceName"`
	DevEUI     string                        `json:"devEUI"`
	Object     map[string]map[string]float64 `json:"object"`
}

type outputInfos struct {
	DeviceEUI string `json:"device_eui"`
	State     string `json:"state"`
	Battery   int    `json:"battery"`
}

type listRequest struct{}

var nc *nats.Conn

func indexHandler(w http.ResponseWriter, r *http.Request) {

	var u deviceInfos

	//	err2 := json.NewDecoder(r.Body).Decode(&u)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err2 := json.Unmarshal(body, &u)

	if err2 != nil {
		http.Error(w, err2.Error(), 400)
		return
	}

	log.Print(u)

	var out outputInfos

	out.DeviceEUI = u.DevEUI
	if p, ok := u.Object["presenceSensor"]; ok {
		for _, v := range p {
			if v > 0.9 {
				out.State = "occupied"
			} else {
				out.State = "free"
			}
			break
		}
	}
	out.Battery = 43

	requestBody, err3 := json.Marshal(out)

	if err3 != nil {
		log.Fatalln(err3)
	}

	log.Printf("sending update %v", out)
	err = nc.Publish("devices.update", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("ok"))
}

func main() {
	var err error

	natsURL, ok := os.LookupEnv("NATS_URL")
	if !ok {
		natsURL = nats.DefaultURL
	}

	options := []nats.Option{}
	tlsConfig, err := loadTLS("NATS")
	if err != nil {
		log.Fatalf("unable to load certificates: %v", err)
	}

	if tlsConfig != nil {
		options = append(options, nats.Secure(tlsConfig))
	}

	nc, err = nats.Connect(natsURL, options...)
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
