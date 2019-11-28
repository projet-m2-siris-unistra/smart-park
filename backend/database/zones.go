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
	var tmp null.String
	var d *string

	err := pool.QueryRowContext(ctx, `
		SELECT zone_id, tenant_id, name, type, color, geo, created_at, updated_at
		FROM zones 
		WHERE zone_id = $1
	`, zoneID).
		Scan(&zone.ZoneID, &zone.TenantID, &zone.Name, &tmp, &zone.Color, &zone.Geography,
			&zone.CreatedAt, &zone.UpdatedAt)

	if err != nil {
		return zone, err
	}

	if tmp.IsZero() == true {
		zone.Type = FreeZone
	} else {
		d = tmp.Ptr()
		switch *d {
		case "paid":
			zone.Type = Paid
		case "blue":
			zone.Type = Blue
		default:
			zone.Type = FreeZone
		}
	}

	return zone, nil
}

// GetZones : get all the zone by the tenant_id
func GetZones(ctx context.Context, tenantID int) ([]Zone, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zones []Zone
	var zone Zone
	var i int
	var tmp null.String
	var d *string

	i = 0
	rows, err := pool.QueryContext(ctx,
		`SELECT z.zone_id, z.tenant_id, z.name, z.type, z.color, z.geo, z.created_at, z.updated_at
		FROM zones z, tenants t
		WHERE z.tenant_id = $1`, tenantID)

	if err != nil {
		return zones, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&zone.ZoneID, &zone.TenantID, &zone.Name, &tmp, &zone.Color, &zone.Geography,
			&zone.CreatedAt, &zone.UpdatedAt)
		if err != nil {
			return zones, err
		}
		if tmp.IsZero() == true {
			zone.Type = FreeZone
		} else {
			d = tmp.Ptr()
			switch *d {
			case "paid":
				zone.Type = Paid
			case "blue":
				zone.Type = Blue
			default:
				zone.Type = FreeZone
			}
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
