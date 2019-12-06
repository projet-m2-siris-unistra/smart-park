package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getDeviceRequest struct {
	DeviceID int `json:"device_id"`
}

type getDevicesRequest struct {
	Limite int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type updateDeviceRequest struct {
	DeviceID int    `json:"device_id"`
	Battery  int    `json:"battery,omitempty"`
	State    string `json:"state,omitempty"`
}

type newDeviceRequest struct {
	Battery int    `json:"battery"`
	State   string `json:"state"`
}

/********************************** GET **********************************/
func getDevice(ctx context.Context, request getDeviceRequest) (database.Device, error) {
	log.Println("handlers: handling getDevice")

	return database.GetDevice(ctx, request.DeviceID)
}

func getFreeDevices(ctx context.Context, request getDevicesRequest) ([]database.Device, error) {
	log.Println("handlers: handling getFreeDevices")

	return database.GetFreeDevices(ctx, request.Limite, request.Offset)
}

func getDevices(ctx context.Context, request getDevicesRequest) ([]database.Device, error) {
	log.Println("handlers: handling getDevices")

	return database.GetDevices(ctx, request.Limite, request.Offset)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/
func updateDevice(ctx context.Context, request updateDeviceRequest) error {
	log.Println("handlers: handling updateDevice")

	err := database.UpdateDevice(ctx, request.DeviceID, request.Battery, request.State)
	return err
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/
func newDevice(ctx context.Context, request newDeviceRequest) error {
	log.Println("handlers: handling newDevice")

	err := database.NewDevice(ctx, request.Battery, request.State)
	return err
}

/********************************** CREATE **********************************/
