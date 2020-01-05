package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

// Value converts a ZoneType to a database/sql/driver.Value
func (s ZoneType) Value() (driver.Value, error) {
	switch s {
	case FreeZone:
		return "free", nil
	case Paid:
		return "paid", nil
	case Blue:
		return "blue", nil
	default:
		return nil, errors.New("invalid ZoneType")
	}
}

// Scan converts a database value to a ZoneType
func (s *ZoneType) Scan(value interface{}) error {
	if sv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := sv.([]byte); ok {
			switch string(v) {
			case "free":
				*s = FreeZone
				return nil
			case "paid":
				*s = Paid
				return nil
			case "blue":
				*s = Blue
				return nil
			}
		}
	}
	return errors.New("failed to scan ZoneType")
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

// ZoneResponse returns the id of the updated / created object
type ZoneResponse struct {
	ZoneID int `json:"zone_id"`
}

/********************************** GET **********************************/

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

// GetZones : get all the zone by the tenant_id
func GetZones(ctx context.Context, tenantID int, paging Paging) ([]Zone, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zones []Zone
	var zone Zone

	rows, err := pool.QueryContext(ctx, fmt.Sprintf(`
		SELECT DISTINCT zone_id, tenant_id, name, type, color, geo, created_at, updated_at
		FROM zones
		WHERE tenant_id = $1
		%s
	`, paging.buildQuery()), tenantID)

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
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return zones, err
	}

	return zones, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdateZone : update a user
func UpdateZone(ctx context.Context, zoneID int, tenantID int,
	name string, zonetype string, color string, geo string) (ZoneResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zone ZoneResponse

	zone.ZoneID = -1

	if (tenantID == 0) && (name == "") && (zonetype == "") && (color == "") && (geo == "") {
		return zone, errors.New("invalid input fields (database/zones.go")
	}

	// modify tenant_id
	if tenantID != 0 {
		err := pool.QueryRowContext(ctx, `
			UPDATE zones SET tenant_id = $1 
			WHERE zone_id = $2 RETURNING zone_id
		`, tenantID, zoneID).Scan(&zone.ZoneID)

		if err == sql.ErrNoRows {
			log.Printf("no zone with id %d\n", zoneID)
			return zone, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return zone, err
		}
	}

	// modify name
	if name != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE zones SET name = $1 
			WHERE zone_id = $2 RETURNING zone_id
		`, name, zoneID).Scan(&zone.ZoneID)

		if err == sql.ErrNoRows {
			log.Printf("no zone with id %d\n", zoneID)
			return zone, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return zone, err
		}
	}

	// modify type
	if zonetype != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE zones SET type = $1 
			WHERE zone_id = $2 RETURNING zone_id
		`, zonetype, zoneID).Scan(&zone.ZoneID)

		if err == sql.ErrNoRows {
			log.Printf("no zone with id %d\n", zoneID)
			return zone, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return zone, err
		}
	}

	// modify color
	if color != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE zones SET color = $1 
			WHERE zone_id = $2 RETURNING zone_id
		`, color, zoneID).Scan(&zone.ZoneID)

		if err == sql.ErrNoRows {
			log.Printf("no zone with id %d\n", zoneID)
			return zone, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return zone, err
		}

	}

	// modify geo
	if geo != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE zones SET geo = $1 
			WHERE zone_id = $2 RETURNING zone_id
		`, geo, zoneID).Scan(&zone.ZoneID)

		if err == sql.ErrNoRows {
			log.Printf("no zone with id %d\n", zoneID)
			return zone, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return zone, err
		}

	}

	return zone, nil
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/

// NewZone : insert a new place
func NewZone(ctx context.Context, tenantID int, name string, zonetype string,
	color string, geo string) (ZoneResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zone ZoneResponse

	zone.ZoneID = -1

	err := pool.QueryRowContext(ctx,
		`INSERT INTO zones
		(
			tenant_id, 
			name,
			type,
			color,
			geo
		) VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING zone_id`, tenantID, name, zonetype, color, geo).Scan(&zone.ZoneID)

	if err == sql.ErrNoRows {
		log.Printf("no zone created\n")
		return zone, err
	}

	if err != nil {
		log.Printf("query error: %v\n", err)
		return zone, err
	}

	return zone, nil
}

/********************************** CREATE **********************************/

/********************************** DELETE **********************************/

// DeleteZone : delete a zone and their places and update all device to 'free'
func DeleteZone(ctx context.Context, zoneID int) (ZoneResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zone ZoneResponse
	var placeID int

	zone.ZoneID = -1

	rows, err := pool.QueryContext(ctx,
		`SELECT DISTINCT place_id FROM places WHERE zone_id = $1`, zoneID)

	if err == sql.ErrNoRows {
		log.Printf("no zone with id %d\n", zoneID)
		return zone, err
	}

	if err != nil {
		log.Printf("query error: %v\n", err)
		return zone, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&placeID)
		if err != nil {
			log.Printf("query error: %v\n", err)
			return zone, err
		}
		_, err = DeletePlace(ctx, placeID)
		if err != nil {
			log.Printf("query error: %v\n", err)
			return zone, err
		}
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Printf("query error: %v\n", err)
		return zone, err
	}

	// update the device id into places
	err = pool.QueryRowContext(ctx, `
			DELETE FROM zones WHERE zone_id = $1 RETURNING zone_id
		`, zoneID).Scan(&zone.ZoneID)

	if err == sql.ErrNoRows {
		log.Printf("no zone with id %d\n", zoneID)
		return zone, err
	}

	if err != nil {
		log.Printf("query error: %v\n", err)
		return zone, err
	}

	return zone, nil
}

/********************************** DELETE **********************************/

/********************************** OPTIONS **********************************/

// CountZone : count number of rows
func CountZone(ctx context.Context, tenantID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var count int

	count = -1

	row := pool.QueryRow("SELECT COUNT(*) FROM zones WHERE tenant_id = $1", tenantID)
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

/********************************** OPTIONS **********************************/
