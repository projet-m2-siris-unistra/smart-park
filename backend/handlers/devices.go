package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getDeviceRequest struct {
	DeviceID int `json:"device_id"`
}

type updateDeviceBatteryRequest struct {
	Battery  int `json:"battery"`
	DeviceID int `json:"device_id"`
}

type updateStateDeviceRequest struct {
	DeviceID int    `json:"device_id"`
	State    string `json:"state"`
}

/********************************** GET **********************************/
func getDevice(ctx context.Context, request getDeviceRequest) (database.Device, error) {
	log.Println("handlers: handling getDevice")

	return database.GetDevice(ctx, request.DeviceID)
}

func getDevices(ctx context.Context, request getDeviceRequest) ([]database.Device, error) {
	log.Println("handlers: handling getDevices")

	return database.GetDevices(ctx)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/
func updateBatteryDevice(ctx context.Context, request updateDeviceBatteryRequest) error {
	log.Println("handlers: handling updateBatteryDevice")

	err := database.UpdateBatteryDevice(ctx, request.DeviceID, request.Battery)
	return err
}

func updateStateDevice(ctx context.Context, request updateStateDeviceRequest) error {
	log.Println("handlers: handling updateStateDevice")

	err := database.UpdateStateDevice(ctx, request.DeviceID, request.State)
	return err
}

/********************************** UPDATE **********************************/
