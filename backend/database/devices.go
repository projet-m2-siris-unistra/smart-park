package database

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"gopkg.in/guregu/null.v3"
)

// DeviceState represents the state of the device
type DeviceState int

const (
	// Free devices have no vehicle on it
	Free DeviceState = iota + 1
	// Occupied devices have a vehicle on it
	Occupied
	// NotAssigned devices
	NotAssigned
)

// MarshalJSON : encode to JSON
func (s DeviceState) MarshalJSON() ([]byte, error) {
	switch s {
	case Free:
		return json.Marshal("free")
	case Occupied:
		return json.Marshal("occupied")
	case NotAssigned:
		return json.Marshal("notassigned")
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
	case "notassigned":
		*s = NotAssigned
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

// UpdateBatteryDevice : update the battery device
func UpdateBatteryDevice(ctx context.Context, deviceID int, battery int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := pool.ExecContext(ctx, `
		UPDATE devices SET battery = $1 
		WHERE device_id = $2
	`, battery, deviceID)

	if err != nil {
		return errors.New("error update device battery")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.New("error : device state - rows affected")
	}

	if rows < 0 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

// UpdateStateDevice : update the state device
func UpdateStateDevice(ctx context.Context, deviceID int, state string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result1 := state == "free"
	result2 := state == "occupied"
	result3 := state == "notassigned"

	if (result1 == false) && (result2 == false) && (result3 == false) {
		return errors.New("invalid device state")
	}

	result, err := pool.ExecContext(ctx, `
		UPDATE devices SET state = $1 
		WHERE device_id = $2
	`, state, deviceID)

	if err != nil {
		return errors.New("error update device state")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errors.New("error : device state - rows affected")
	}

	if rows < 0 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

// GetDevice fetches the device by its ID
func GetDevice(ctx context.Context, deviceID int) (Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var device Device
	var tmp null.String
	var d *string

	err := pool.QueryRowContext(ctx, `
		SELECT device_id, battery, state, created_at, updated_at
		FROM devices 
		WHERE device_id = $1
	`, deviceID).
		Scan(&device.DeviceID, &device.Battery, &tmp,
			&device.CreatedAt, &device.UpdatedAt)

	if err != nil {
		return device, err
	}

	if tmp.IsZero() == true {
		device.State = NotAssigned
	} else {
		d = tmp.Ptr()
		switch *d {
		case "free":
			device.State = Free
		case "occupied":
			device.State = Occupied
		default:
			device.State = NotAssigned
		}
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
	var d *string
	var tmp null.String

	i = 0
	rows, err := pool.QueryContext(ctx,
		`SELECT device_id, battery, state, created_at, updated_at FROM devices`)

	if err != nil {
		return devices, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&device.DeviceID, &device.Battery, &tmp,
			&device.CreatedAt, &device.UpdatedAt)
		if err != nil {
			return devices, err
		}
		if tmp.IsZero() == true {
			device.State = NotAssigned
		} else {
			d = tmp.Ptr()
			switch *d {
			case "free":
				device.State = Free
			case "occupied":
				device.State = Occupied
			default:
				device.State = NotAssigned
			}
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
