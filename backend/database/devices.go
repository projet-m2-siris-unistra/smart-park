package database

import (
	"context"
	"time"
)

// DeviceState represents the state of the device
type DeviceState int

const (
	// Free devices have no vehicle on it
	Free DeviceState = iota + 1
	// Occupied devices have a vehicle on it
	Occupied
)

// Device represents an IoT device
type Device struct {
	DeviceID  int       `json:"device_id"`
	Battery   int       `json:"battery"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetDevice fetches the device by its ID
func GetDevice(ctx context.Context, deviceID int) (Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var device Device

	err := pool.QueryRowContext(ctx, `
		SELECT device_id, battery, state, created_at, updated_at
		FROM devices 
		WHERE device_id = $1
	`, deviceID).
		Scan(&device.DeviceID, &device.Battery, &device.State,
			&device.CreatedAt, &device.UpdatedAt)

	if err != nil {
		return device, err
	}

	return device, nil
}
