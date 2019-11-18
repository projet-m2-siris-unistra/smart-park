package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getDeviceRequest struct {
	DeviceID int `json:"device_id"`
}

func getDevice(m *nats.Msg) {
	ctx := context.TODO()
	log.Println("handlers: handling getDevice")

	var request getDeviceRequest
	err := json.Unmarshal(m.Data, &request)
	if err != nil {
		log.Println(err)
		return
	}

	device, err := database.GetDevice(ctx, request.DeviceID)
	if err != nil {
		log.Println(err)
		return
	}

	payload, err := json.Marshal(device)
	if err != nil {
		log.Println(err)
		return
	}

	m.Respond(payload)
}
