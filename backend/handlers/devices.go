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
	TenantID int `json:"tenant_id,omitempty"`
}

type updateDeviceRequest struct {
	DeviceID int    `json:"device_id"`
	Battery  int    `json:"battery,omitempty"`
	State    string `json:"state,omitempty"`
	TenantID int 		`json:"tenant_id,omitempty"`
	DeviceEUI string 	`json:"device_eui,omitempty"`
}

type newDeviceRequest struct {
	Battery int    `json:"battery"`
	State   string `json:"state"`
	TenantID int 		`json:"tenant_id"`
	DeviceEUI string 	`json:"device_eui"`
}

type resultListDevice struct {
	Count int `json:"count"`
	Data []database.Device `json:"data"`
}

/********************************** GET **********************************/
func getDevice(ctx context.Context, request getDeviceRequest) (database.Device, error) {
	log.Println("handlers: handling getDevice")

	return database.GetDevice(ctx, request.DeviceID)
}

func getFreeDevices(ctx context.Context, request getDevicesRequest) (resultListDevice, error) {
	log.Println("handlers: handling getFreeDevices")

	var result resultListDevice
	var err error 
	result.Count, err = database.CountDeviceFree(ctx, request.TenantID)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetFreeDevices(ctx, request.Limite, request.Offset, request.TenantID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getNotAssignedDevices(ctx context.Context, request getDevicesRequest) (resultListDevice, error) {
	log.Println("handlers: handling getNotAssignedDevices")

	var result resultListDevice
	var err error 
	result.Count, err = database.CountDeviceNotAssigned(ctx, request.TenantID)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetNotAssignedDevices(ctx, request.Limite, request.Offset, request.TenantID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getDevices(ctx context.Context, request getDevicesRequest) (resultListDevice, error) {
	log.Println("handlers: handling getDevices")

	var result resultListDevice
	var err error 
	result.Count, err = database.CountDevice(ctx)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetDevices(ctx, request.Limite, request.Offset)
	if err != nil {
		return result, err
	}
	return result, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/
func updateDevice(ctx context.Context, request updateDeviceRequest) (database.DeviceResponse, error) {
	log.Println("handlers: handling updateDevice")

	return database.UpdateDevice(ctx, request.DeviceID, request.Battery, request.State, request.TenantID, 
		request.DeviceEUI)
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/
func newDevice(ctx context.Context, request newDeviceRequest) (database.DeviceResponse, error) {
	log.Println("handlers: handling newDevice")

	return database.NewDevice(ctx, request.Battery, request.State, request.TenantID, 
		request.DeviceEUI)
}

/********************************** CREATE **********************************/

/********************************** DELETE **********************************/

func deleteDevice(ctx context.Context, request getDeviceRequest) (database.DeviceResponse, error) {
	log.Println("handlers: handling deleteDevice")

	return database.DeleteDevice(ctx, request.DeviceID)
}

/********************************** DELETE **********************************/