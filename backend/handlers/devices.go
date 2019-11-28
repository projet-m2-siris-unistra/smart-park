package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getDeviceRequest struct {
	DeviceID int `json:"device_id"`
}

type updateDeviceRequest struct {
	DeviceID int    `json:"device_id"`
	Battery  int    `json:"battery,omitempty"`
	State    string `json:"state,omitempty"`
}

/********************************** GET **********************************/
func getDevice(ctx context.Context, request getDeviceRequest) (database.Device, error) {
	log.Println("handlers: handling getDevice")

	return database.GetDevice(ctx, request.DeviceID)
}

func getFreeDevices(ctx context.Context, request getDeviceRequest) ([]database.Device, error) {
	log.Println("handlers: handling getFreeDevices")

	return database.GetFreeDevices(ctx)
}

func getDevices(ctx context.Context, request getDeviceRequest) ([]database.Device, error) {
	log.Println("handlers: handling getDevices")

	return database.GetDevices(ctx)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/
func updateDevice(ctx context.Context, request updateDeviceRequest) error {
	log.Println("handlers: handling updateDevice")

	err := database.UpdateDevice(ctx, request.DeviceID, request.Battery, request.State)
	return err
}

/********************************** UPDATE **********************************/
