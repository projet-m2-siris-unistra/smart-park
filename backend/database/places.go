package database

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v3"
)

// Place : owner of one park
type Place struct {
	PlaceID   int         `json:"place_id"`
	ZoneID    int         `json:"zone_id"`
	Type      string      `json:"type"`
	Geography null.String `json:"geo"`
	DeviceID  int         `json:"device_id"`
	Timestamps
}

// GetPlace fetches the place by its ID
func GetPlace(ctx context.Context, placeID int) (Place, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var place Place

	err := pool.QueryRowContext(ctx, `
		SELECT place_id, zone_id, type, geo, place_id, created_at, updated_at
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

// GetPlaces : get all the place
func GetPlaces(ctx context.Context, zoneID int) ([]Place, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var places []Place
	var place Place
	var i int

	i = 0
	rows, err := pool.QueryContext(ctx,
		`SELECT place_id, zone_id, type, geo, place_id, created_at, updated_at
		FROM places WHERE zone_id = $1`, zoneID)

	if err != nil {
		return places, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&place.PlaceID, &place.ZoneID, &place.Type, &place.Geography, &place.DeviceID,
			&place.CreatedAt, &place.UpdatedAt)
		if err != nil {
			return places, err
		}
		places = append(places, place)
		i = i + 1
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return places, err
	}

	return places, nil
}
