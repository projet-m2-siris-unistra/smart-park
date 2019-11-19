package database

import (
	"context"
	"encoding/json"
	"errors"
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

// MarshalJSON : encode to JSON
func (s DeviceState) MarshalJSON() ([]byte, error) {
	switch s {
	case Free:
		return json.Marshal("free")
	case Occupied:
		return json.Marshal("occupied")
	}

	return nil, errors.New("invalid device state")
}

// UnmarshalJSON : decode JSON
func (s *DeviceState) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	switch j {
	case "free":
		*s = Free
	case "occupied":
		*s = Occupied
	default:
		return errors.New("invalid DeviceState")
	}

	return nil
}

// Device represents an IoT device
type Device struct {
	DeviceID int         `json:"device_id"`
	Battery  int         `json:"battery"`
	State    DeviceState `json:"state"`
	Timestamps
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

// GetDevices : get all the device
func GetDevices(ctx context.Context) ([]Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var devices []Device
	var device Device
	var i int

	i = 0
	rows, err := pool.QueryContext(ctx,
		`SELECT device_id, battery, state, created_at, updated_at FROM devices`)

	if err != nil {
		return devices, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&device.DeviceID, &device.Battery, &device.State,
			&device.CreatedAt, &device.UpdatedAt)
		if err != nil {
			return devices, err
		}
		devices = append(devices, device)
		i = i + 1
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return devices, err
	}

	return devices, nil
}
