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
	database.Paging
	TenantID int `json:"tenant_id,omitempty"`
}

type updateDeviceRequest struct {
	DeviceID  *int    `json:"device_id,omitempty"`
	Battery   *int    `json:"battery,omitempty"`
	State     *string `json:"state,omitempty"`
	TenantID  *int    `json:"tenant_id,omitempty"`
	DeviceEUI *string `json:"device_eui,omitempty"`
}

type newDeviceRequest struct {
	Battery   int    `json:"battery"`
	State     string `json:"state"`
	TenantID  int    `json:"tenant_id"`
	DeviceEUI string `json:"device_eui"`
}

type resultListDevice struct {
	Count int               `json:"count"`
	Data  []database.Device `json:"data"`
}

/********************************** GET **********************************/
func getDevice(ctx context.Context, request getDeviceRequest) (database.Device, error) {
	log.Println("handlers: handling getDevice")

	return database.GetDevice(ctx, request.DeviceID)
}

func getFreeDevices(ctx context.Context, request getDevicesRequest) (resultListDevice, error) {
	log.Println("handlers: handling getFreeDevices")

	filter := database.DeviceFilter{}.
		WithTenantID(request.TenantID).
		WithState(database.Free)

	var result resultListDevice
	var err error
	result.Count, err = database.CountDevices(ctx, filter)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetDevices(ctx, filter, request.Paging)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getNotAssignedDevices(ctx context.Context, request getDevicesRequest) (resultListDevice, error) {
	log.Println("handlers: handling getNotAssignedDevices")

	filter := database.DeviceFilter{}.
		WithTenantID(request.TenantID).
		WithIsAttached(false)

	var result resultListDevice
	var err error
	result.Count, err = database.CountDevices(ctx, filter)
	if err != nil {
		return result, err
	}

	result.Data, err = database.GetDevices(ctx, filter, request.Paging)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getDevices(ctx context.Context, request getDevicesRequest) (resultListDevice, error) {
	log.Println("handlers: handling getDevices")

	filter := database.DeviceFilter{}.
		WithTenantID(request.TenantID)

	var result resultListDevice
	var err error
	result.Count, err = database.CountDevices(ctx, filter)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetDevices(ctx, filter, request.Paging)
	if err != nil {
		return result, err
	}
	return result, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/
func updateDevice(ctx context.Context, request updateDeviceRequest) (bool, error) {
	log.Println("handlers: handling updateDevice")

	err := database.UpdateDevice(ctx, request.DeviceID, request.Battery, request.State, request.TenantID, request.DeviceEUI)
	if err != nil {
		return false, err
	}
	return true, nil
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

func deleteDevice(ctx context.Context, request getDeviceRequest) (bool, error) {
	log.Println("handlers: handling deleteDevice")

	return database.DeleteDevice(ctx, request.DeviceID)
}

/********************************** DELETE **********************************/
