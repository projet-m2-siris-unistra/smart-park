package database

import (
	"context"
	"time"
)

// Zone :
type Zone struct {
	ZoneID    int       `json:"zone_id"`
	TenantID  int       `json:"tenant_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Color     string    `json:"color"`
	Geography string    `json:"geo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
