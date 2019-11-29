package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

// sensor structure

type infos struct {
	ID      int    `json:"device_id"`
	State   string `json:"state"`
	Battery int    `json:"battery"`
}

// return true or false

func rand2() bool {
	return rand.Int31()&0x01 == 0
}

var nc *nats.Conn

// message flood

func flood(id int) {
	for {
		bat := rand.Intn(1500)
		occ := rand2()

		state := "free"
		if occ {
			state = "occupied"
		}

		test := infos{
			ID:      id,
			State:   state,
			Battery: bat,
		}

		requestBody, err := json.Marshal(test)

		if err != nil {
			log.Fatalln(err)
		}

		err = nc.Publish("devices.update.battery", requestBody)
		if err != nil {
			log.Fatal(err)
		}
		err = nc.Publish("devices.update.state", requestBody)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Minute)
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

	ids := os.Getenv("DEVICE_IDS")
	for _, id := range strings.Split(ids, ",") {
		i, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		go flood(i)
	}

	// handle SIGINT signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// close connection
	nc.Close()
}
