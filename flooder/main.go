package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"
)

type deviceInfos struct {
	ID      int    `json:"device_id"`
	State   string `json:"state"`
	Battery int    `json:"battery"`
}

type listRequest struct{}
type listResponse []deviceInfos

func rand2() bool {
	return rand.Int31()&0x01 == 0
}

var nc *nats.Conn

func flood(id int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(60)))
	for {
		bat := rand.Intn(100)
		occ := rand2()

		state := "free"
		if occ {
			state = "occupied"
		}

		payload := deviceInfos{
			ID:      id,
			State:   state,
			Battery: bat,
		}

		requestBody, err := json.Marshal(payload)

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("sending update %v", payload)
		err = nc.Publish("devices.update", requestBody)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Minute + time.Second*time.Duration(rand.Intn(5)))
	}
}

func main() {
	var err error

	natsURL, ok := os.LookupEnv("NATS_URL")
	if !ok {
		natsURL = nats.DefaultURL
	}

	nc, err = nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}

	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	var resp listResponse
	req := listRequest{}
	err = c.Request("devices.list", req, &resp, 100*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range resp {
		log.Printf("starting to flood %v", device)
		go flood(device.ID)
	}

	// handle SIGINT signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	// close connection
	c.Close()
}
