package database

import (
	"context"
	"time"
)

// GetStatsDevice :
func GetStatsDevice(ctx context.Context) ([]int, []int, []int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var devices []int
	var stateDevices []int
	var zones []int

	var device int
	var stateDevice string
	var zone int

	rows, err := pool.QueryContext(ctx,
		`SELECT d.tenant_id, d.state, d.device_id FROM devices d FULL JOIN
		 tenants t on t.tenant_id=d.device_id`)

	if err != nil {
		return devices, stateDevices, zones, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&device, &stateDevice, &zone)
		if err != nil {
			return devices, stateDevices, zones, err
		}

		devices = append(devices, device)
		if stateDevice == "free" {
			stateDevices = append(stateDevices, 1)
		} else {
			stateDevices = append(stateDevices, 0)
		}
		zones = append(zones, zone)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return devices, stateDevices, zones, err
	}

	return devices, stateDevices, zones, nil
}
