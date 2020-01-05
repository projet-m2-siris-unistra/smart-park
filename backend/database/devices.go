package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
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

// Value converts a DeviceState to a database/sql/driver.Value
func (s DeviceState) Value() (driver.Value, error) {
	switch s {
	case Free:
		return "free", nil
	case Occupied:
		return "occupied", nil
	case NotAssigned:
		return nil, nil
	default:
		return nil, errors.New("invalid ZoneType")
	}
}

// Scan converts a database value to a DeviceState
func (s *DeviceState) Scan(value interface{}) error {
	if value == nil {
		*s = NotAssigned
		return nil
	}

	if sv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := sv.([]byte); ok {
			switch string(v) {
			case "free":
				*s = Free
				return nil
			case "occupied":
				*s = Occupied
				return nil
			}
		}
	}

	return errors.New("failed to scan DeviceState")
}

// Device represents an IoT device
type Device struct {
	DeviceID  int         `json:"device_id"`
	Battery   int         `json:"battery"`
	State     DeviceState `json:"state"`
	TenantID  int         `json:"tenant_id"`
	DeviceEUI string      `json:"device_eui"`
	Timestamps
}

// DeviceResponse returns the id of the updated / created object
type DeviceResponse struct {
	DeviceID int `json:"device_id"`
}

/********************************** GET **********************************/

// GetDevice fetches the device by its ID
func GetDevice(ctx context.Context, deviceID int) (Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var device Device

	err := pool.QueryRowContext(ctx, `
		SELECT device_id, battery, state, tenant_id, device_eui, created_at, updated_at
		FROM devices 
		WHERE device_id = $1
	`, deviceID).
		Scan(&device.DeviceID, &device.Battery, &device.State, &device.TenantID, &device.DeviceEUI,
			&device.CreatedAt, &device.UpdatedAt)

	if err != nil {
		return device, err
	}

	return device, nil
}

// GetDevices : get all the device
func GetDevices(ctx context.Context, limite int, offset int) ([]Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var devices []Device
	var device Device

	limite, offset = CheckArgDevice(limite, offset)

	rows, err := pool.QueryContext(ctx,
		`SELECT DISTINCT device_id, battery, state, tenant_id, device_eui, created_at, updated_at FROM devices
		LIMIT $1 OFFSET $2`, limite, offset)

	if err != nil {
		return devices, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&device.DeviceID, &device.Battery, &device.State, &device.TenantID, &device.DeviceEUI,
			&device.CreatedAt, &device.UpdatedAt)
		if err != nil {
			return devices, err
		}
		devices = append(devices, device)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return devices, err
	}

	return devices, nil
}

// GetFreeDevices : get all the free devices
func GetFreeDevices(ctx context.Context, limite int, offset int, tenantID int) ([]Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var devices []Device
	var device Device

	limite, offset = CheckArgDevice(limite, offset)

	rows, err := pool.QueryContext(ctx,
		`SELECT DISTINCT device_id, battery, state, tenant_id, device_eui, created_at, updated_at FROM devices
		WHERE state='free' AND tenant_id = $1 LIMIT $2 OFFSET $3`, tenantID, limite, offset)

	if err != nil {
		return devices, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&device.DeviceID, &device.Battery, &device.State, &device.TenantID, &device.DeviceEUI,
			&device.CreatedAt, &device.UpdatedAt)
		if err != nil {
			return devices, err
		}
		devices = append(devices, device)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return devices, err
	}

	return devices, nil
}

// GetNotAssignedDevices : get all the not assigned devices
func GetNotAssignedDevices(ctx context.Context, limite int, offset int, tenantID int) ([]Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var devices []Device
	var device Device

	limite, offset = CheckArgDevice(limite, offset)

	rows, err := pool.QueryContext(ctx,
		`SELECT device_id, battery, state, tenant_id, device_eui, created_at, updated_at FROM devices 
		WHERE tenant_id = $1 AND device_id NOT IN (
			SELECT DISTINCT device_id FROM places WHERE device_id is not null
			) 
		LIMIT $2 OFFSET $3`, tenantID, limite, offset)

	if err != nil {
		return devices, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&device.DeviceID, &device.Battery, &device.State, &device.TenantID, &device.DeviceEUI,
			&device.CreatedAt, &device.UpdatedAt)
		if err != nil {
			return devices, err
		}
		devices = append(devices, device)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return devices, err
	}

	return devices, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdateDevice : update a device - need the device ID
func UpdateDevice(ctx context.Context, deviceID int, battery int, state string, tenantID int,
	deviceEUI string) (DeviceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var device DeviceResponse

	device.DeviceID = -1

	if (state == "") && (battery == 0) && (tenantID == 0) && (deviceEUI == "") {
		return device, errors.New("invalid input fields (database/devices.go")
	}

	if battery != 0 {

		err := pool.QueryRowContext(ctx, `
			UPDATE devices SET battery = $1 
			WHERE device_id = $2 RETURNING device_id
		`, battery, deviceID).Scan(&device.DeviceID)

		if err == sql.ErrNoRows {
			log.Printf("no device with id %d\n", deviceID)
			return device, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return device, err
		}
	}

	if state != "" {

		err := pool.QueryRowContext(ctx, `
			UPDATE devices SET state = $1 
			WHERE device_id = $2 RETURNING device_id
		`, state, deviceID).Scan(&device.DeviceID)

		if err == sql.ErrNoRows {
			log.Printf("no device with id %d\n", deviceID)
			return device, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return device, err
		}
	}

	if tenantID != 0 {

		err := pool.QueryRowContext(ctx, `
			UPDATE devices SET tenant_id = $1 
			WHERE device_id = $2 RETURNING device_id
		`, tenantID, deviceID).Scan(&device.DeviceID)

		if err == sql.ErrNoRows {
			log.Printf("no device with id %d\n", deviceID)
			return device, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return device, err
		}
	}

	if deviceEUI != "" {

		err := pool.QueryRowContext(ctx, `
			UPDATE devices SET device_eui = $1 
			WHERE device_id = $2 RETURNING device_id
		`, deviceEUI, deviceID).Scan(&device.DeviceID)

		if err == sql.ErrNoRows {
			log.Printf("no device with id %d\n", deviceID)
			return device, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return device, err
		}
	}

	return device, nil
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/

// NewDevice : insert a new device
func NewDevice(ctx context.Context, battery int, state string, tenantID int,
	deviceEUI string) (DeviceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var device DeviceResponse

	device.DeviceID = -1

	err := pool.QueryRowContext(ctx,
		`INSERT INTO 
		devices (
			battery, 
			state,
			tenant_id,
			device_eui
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) RETURNING device_id`, battery, state, tenantID, deviceEUI).Scan(&device.DeviceID)

	if err == sql.ErrNoRows {
		log.Printf("no device created\n")
		return device, err
	}

	if err != nil {
		log.Printf("query error: %v\n", err)
		return device, err
	}

	return device, nil
}

/********************************** CREATE **********************************/

/********************************** DELETE **********************************/

// DeleteDevice : delete a device
func DeleteDevice(ctx context.Context, deviceID int) (DeviceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var device DeviceResponse

	device.DeviceID = -1

	// verify if the device is assigned or not
	err := pool.QueryRowContext(ctx, `
		SELECT DISTINCT device_id
		FROM devices WHERE device_id = $1 AND device_id NOT IN (
			SELECT DISTINCT device_id FROM places WHERE device_id is not null
		) 
	`, deviceID).Scan(&device.DeviceID)

	// the device is assigned
	if err != nil {
		log.Printf("device_id %d is assigned, delete impossible\n", deviceID)
		return device, err
	}

	// delete the device
	err = pool.QueryRowContext(ctx, `
		DELETE FROM devices WHERE device_id = $1 RETURNING device_id
	`, deviceID).Scan(&device.DeviceID)

	if err == sql.ErrNoRows {
		log.Printf("no device with id %d\n", deviceID)
		return device, err
	}

	if err != nil {
		log.Printf("query error: %v\n", err)
		return device, err
	}

	return device, nil
}

/********************************** DELETE **********************************/

/********************************** OPTIONS **********************************/

// CountDeviceFree : count number of rows
func CountDeviceFree(ctx context.Context, tenantID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var count int

	count = -1

	row := pool.QueryRow("select count(*) from devices where tenant_id = $1 and state='free'", tenantID)
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// CountDeviceNotAssigned : count number of rows
func CountDeviceNotAssigned(ctx context.Context, tenantID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var count int

	count = -1

	row := pool.QueryRow("select count(*) from devices where tenant_id = $1 and device_id not in (select distinct device_id from places where device_id is not null)", tenantID)
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// CountDevice : count number of rows
func CountDevice(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var count int

	count = -1

	row := pool.QueryRow("SELECT COUNT(*) FROM devices")
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// CheckArgDevice : check limit and offset arguments
func CheckArgDevice(limite int, offset int) (int, int) {

	if limite == 0 {
		limite = 20
	}

	if offset == 0 {
		offset = 0
	}

	return limite, offset
}

/********************************** OPTIONS **********************************/
