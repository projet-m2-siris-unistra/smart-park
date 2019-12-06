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

/********************************** GET **********************************/

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
func GetDevices(ctx context.Context, limite int, offset int) ([]Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var devices []Device
	var device Device
	var i int
	var d *string
	var tmp null.String

	i = 0
	if (limite != 0 && offset != 0) {
		rows, err := pool.QueryContext(ctx,
		`SELECT device_id, battery, state, created_at, updated_at FROM devices
		LIMIT $1 OFFSET $2`, limite, offset)

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
	} else if (limite != 0) {
		rows, err := pool.QueryContext(ctx,
			`SELECT device_id, battery, state, created_at, updated_at FROM devices
			LIMIT $1`, limite)
	
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
	} else if (offset != 0) {
		rows, err := pool.QueryContext(ctx,
			`SELECT device_id, battery, state, created_at, updated_at FROM devices
			OFFSET $1`, offset)
	
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
	} else {
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
	}

	return devices, nil
}

// GetFreeDevices : get all the avalaible devices
func GetFreeDevices(ctx context.Context, limite int, offset int) ([]Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var devices []Device
	var device Device
	var i int
	var d *string
	var tmp null.String

	i = 0
	if (limite != 0 && offset != 0) {
		rows, err := pool.QueryContext(ctx,
		`SELECT device_id, battery, state, created_at, updated_at FROM devices
		WHERE state='free' LIMIT $1 OFFSET $2`, limite, offset)

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
	} else if (limite != 0) {
		rows, err := pool.QueryContext(ctx,
			`SELECT device_id, battery, state, created_at, updated_at FROM devices
			WHERE state='free' LIMIT $1`, limite)
	
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
	} else if (offset != 0) {
		rows, err := pool.QueryContext(ctx,
			`SELECT device_id, battery, state, created_at, updated_at FROM devices
			WHERE state='free' OFFSET $1`, offset)
	
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
	} else {
		rows, err := pool.QueryContext(ctx,
			`SELECT device_id, battery, state, created_at, updated_at FROM devices
			WHERE state='free'`)
	
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
	}
	return devices, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdateDevice : update a device - need the device ID
func UpdateDevice(ctx context.Context, deviceID int, battery int, state string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if (state == "") && (battery == 0) {
		return errors.New("invalid input fields (database/devices.go")
	}

	// update on the battery only
	if (state == "") && (battery != 0) {

		result, err := pool.ExecContext(ctx, `
			UPDATE devices SET battery = $1 
			WHERE device_id = $2
		`, battery, deviceID)

		if err != nil {
			return errors.New("error update device battery")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : device state - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}

		// update the state only
	} else if (state != "") && (battery == 0) {
		// have to verify if the input is correctly write
		if ((state == "free") == false) && ((state == "occupied") == false) && ((state == "notassigned") == false) {
			return errors.New("invalid device state")
		}

		result, err := pool.ExecContext(ctx, `
			UPDATE devices SET state = $1 
			WHERE device_id = $2
		`, state, deviceID)

		if err != nil {
			return errors.New("error update device state")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : device state - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}

		// update the battery and the state
	} else {

		// have to verify if the input is correctly write
		if ((state == "free") == false) && ((state == "occupied") == false) && ((state == "notassigned") == false) {
			return errors.New("invalid device state")
		}
		result, err := pool.ExecContext(ctx, `
			UPDATE devices SET battery = $1, state = $2
			WHERE device_id = $3
		`, battery, state, deviceID)

		if err != nil {
			return errors.New("error update device battery")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : device state - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}

	}
	return nil
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/

// NewDevice : insert a new device
func NewDevice(ctx context.Context, battery int, state string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := pool.ExecContext(ctx,
		`INSERT INTO 
		devices (
			battery, 
			state
		) VALUES (
			$1,
			$2
		)`, battery, state)

	if err != nil {
		return errors.New("error new device")
	}

	return nil
}

/********************************** CREATE **********************************/
