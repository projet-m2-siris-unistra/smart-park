package database

import (
	"context"
	"time"
)

// Place : owner of one park
type Place struct {
	PlaceID   int       `json:"place_id"`
	ZoneID    int       `json:"zone_id"`
	Type      string    `json:"type"`
	Geography string    `json:"geo"`
	DeviceID  int       `json:"device_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetPlace fetches the place by its ID
func GetPlace(ctx context.Context, placeID int) (Place, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var place Place

	err := pool.QueryRowContext(ctx, `
		SELECT place_id, zone_id, type, geo, device_id, created_at, updated_at
		FROM places 
		WHERE place_id = $1
	`, placeID).
		Scan(&place.PlaceID, &place.ZoneID, &place.Type, &place.Geography, &place.DeviceID,
			&place.CreatedAt, &place.UpdatedAt)

	if err != nil {
		return place, err
	}

	return place, nil
}
