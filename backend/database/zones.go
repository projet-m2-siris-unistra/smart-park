package database

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"gopkg.in/guregu/null.v3"
)

// ZoneType represents the type of the zone
type ZoneType int

const (
	// FreeZone zones
	FreeZone ZoneType = iota + 1
	// Paid zones
	Paid
	// Blue zones
	Blue
)

// MarshalJSON : encode to JSON
func (s ZoneType) MarshalJSON() ([]byte, error) {
	switch s {
	case FreeZone:
		return json.Marshal("free")
	case Paid:
		return json.Marshal("paid")
	case Blue:
		return json.Marshal("blue")
	}

	return nil, errors.New("invalid zone type")
}

// UnmarshalJSON : decode JSON
func (s *ZoneType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	switch j {
	case "free":
		*s = FreeZone
	case "paid":
		*s = Paid
	case "blue":
		*s = Blue
	default:
		return errors.New("invalid ZoneType")
	}

	return nil
}

// Zone :
type Zone struct {
	ZoneID    int         `json:"zone_id"`
	TenantID  int         `json:"tenant_id"`
	Name      string      `json:"name"`
	Type      ZoneType    `json:"type"`
	Color     null.String `json:"color"`
	Geography null.String `json:"geo"`
	Timestamps
}

// GetZone fetches the zone by its ID
func GetZone(ctx context.Context, zoneID int) (Zone, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zone Zone

	err := pool.QueryRowContext(ctx, `
		SELECT zone_id, tenant_id, name, type, color, geo, created_at, updated_at
		FROM zones 
		WHERE zone_id = $1
	`, zoneID).
		Scan(&zone.ZoneID, &zone.TenantID, &zone.Name, &zone.Type, &zone.Color, &zone.Geography,
			&zone.CreatedAt, &zone.UpdatedAt)

	if err != nil {
		return zone, err
	}

	return zone, nil
}

// GetZones : get all the zone
func GetZones(ctx context.Context) ([]Zone, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zones []Zone
	var zone Zone
	var i int

	i = 0
	rows, err := pool.QueryContext(ctx,
		`SELECT zone_id, tenant_id, name, type, color, geo, created_at, updated_at
		FROM zones `)

	if err != nil {
		return zones, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&zone.ZoneID, &zone.TenantID, &zone.Name, &zone.Type, &zone.Color, &zone.Geography,
			&zone.CreatedAt, &zone.UpdatedAt)
		if err != nil {
			return zones, err
		}
		zones = append(zones, zone)
		i = i + 1
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return zones, err
	}

	return zones, nil
}
