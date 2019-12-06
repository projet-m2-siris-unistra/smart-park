package database

import (
	"context"
	"encoding/json"
	"errors"
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

/********************************** GET **********************************/

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
func GetZones(ctx context.Context, tenantID int, limite int, offset int) ([]Zone, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var zones []Zone
	var zone Zone
	var i int
	var tmp null.String
	var d *string

	i = 0
	if (limite != 0 && offset != 0) {
		rows, err := pool.QueryContext(ctx,
		`SELECT z.zone_id, z.tenant_id, z.name, z.type, z.color, z.geo, z.created_at, z.updated_at
		FROM zones z, tenants t
		WHERE z.tenant_id = $1 LIMIT $2 OFFSET $3`, tenantID, limite, offset)

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
	} else if (limite != 0) {
		rows, err := pool.QueryContext(ctx,
			`SELECT z.zone_id, z.tenant_id, z.name, z.type, z.color, z.geo, z.created_at, z.updated_at
			FROM zones z, tenants t
			WHERE z.tenant_id = $1 LIMIT $2`, tenantID, limite)
	
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
	} else if (offset != 0) {
		rows, err := pool.QueryContext(ctx,
			`SELECT z.zone_id, z.tenant_id, z.name, z.type, z.color, z.geo, z.created_at, z.updated_at
			FROM zones z, tenants t
			WHERE z.tenant_id = $1 OFFSET $2`, tenantID, offset)
	
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
	} else {
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
	}

	return zones, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdateZone : update a user
func UpdateZone(ctx context.Context, zoneID int, tenantID int,
	name string, zonetype string, color string, geo string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if (tenantID == 0) && (name == "") && (zonetype == "") && (color == "") && (geo == "") {
		return errors.New("invalid input fields (database/zones.go")
	}

	// modify tenant_id
	if tenantID != 0 {
		result, err := pool.ExecContext(ctx, `
			UPDATE zones SET tenant_id = $1 
			WHERE zone_id = $2
		`, tenantID, zoneID)

		if err != nil {
			return errors.New("error update zone tenant_id")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : zone tenant_id - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify name
	if name != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE zones SET name = $1 
			WHERE zone_id = $2
		`, name, zoneID)

		if err != nil {
			return errors.New("error update zone name")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : zone name - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify type
	if zonetype != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE zones SET type = $1 
			WHERE zone_id = $2
		`, zonetype, zoneID)

		if err != nil {
			return errors.New("error update zone type")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : zone type - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify color
	if color != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE zones SET color = $1 
			WHERE zone_id = $2
		`, color, zoneID)

		if err != nil {
			return errors.New("error update zone color")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : zone color - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify geo
	if geo != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE zones SET geo = $1 
			WHERE zone_id = $2
		`, geo, zoneID)

		if err != nil {
			return errors.New("error update zone geo")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : zone geo - rows affected")
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

// NewZone : insert a new place
func NewZone(ctx context.Context, tenantID int, name string, zonetype string,
	color string, geo string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := pool.ExecContext(ctx,
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
		)`, tenantID, name, zonetype, color, geo)

	if err != nil {
		return errors.New("error new place")
	}

	return nil
}

/********************************** CREATE **********************************/
