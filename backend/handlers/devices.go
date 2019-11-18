package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getDeviceRequest struct {
	DeviceID int `json:"device_id"`
}

func getDevice(ctx context.Context, request getDeviceRequest) (database.Device, error) {
	log.Println("handlers: handling getDevice")

	return database.GetDevice(ctx, request.DeviceID)
}
