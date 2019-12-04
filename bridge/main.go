package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"
)

type deviceInfos struct {
	DeviceName      string `json:"deviceName"`
	DevEUI          string `json:"devEUI"`
	ApplicationName string `json:"applicationName"`
}

type outputInfos struct {
	Device_EUI		string `json:"device_EUI"`
	State			string `json:"state"`
	Battery			int    `json:"battery"`
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

	out.Device_EUI = u.DevEUI
	out.State = "occupied"
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

	//w.Write([]byte("DeviceName : %s, DevEUI : %s, Battery : %s\n",u.deviceName, u.devEUI, u.applicationName))


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
