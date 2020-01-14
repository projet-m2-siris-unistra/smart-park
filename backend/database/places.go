package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"gopkg.in/guregu/null.v3"
)

// Place : owner of one park
type Place struct {
	PlaceID   int         `json:"place_id"`
	ZoneID    int         `json:"zone_id"`
	Type      string      `json:"type"`
	Geography null.String `json:"geo"`
	DeviceID  null.Int    `json:"device_id"`
	Timestamps
}

// PlaceResponse returns the id of the updated / created object
type PlaceResponse struct {
	PlaceID int `json:"place_id"`
}

/********************************** GET **********************************/

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

// GetPlaces : get all the place
func GetPlaces(ctx context.Context, zoneID int, paging Paging) ([]Place, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var places []Place
	var place Place

	rows, err := pool.QueryContext(
		ctx,
		fmt.Sprintf(`
			SELECT DISTINCT place_id, zone_id, type, geo, device_id, created_at, updated_at
			FROM places 
			WHERE zone_id = $1
			%s
		`, paging.buildQuery()),
		zoneID,
	)

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
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return places, err
	}

	return places, nil
}

// GetPlacesWithNoDevice : get all place with a null device_id which means the place has not an assigned device
func GetPlacesWithNoDevice(ctx context.Context) ([]Place, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var places []Place
	var place Place

	rows, err := pool.QueryContext(ctx,
		`SELECT DISTINCT place_id, zone_id, type, geo, created_at, updated_at
		from places where device_id is null`)

	if err != nil {
		return places, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&place.PlaceID, &place.ZoneID, &place.Type, &place.Geography,
			&place.CreatedAt, &place.UpdatedAt)
		if err != nil {
			return places, err
		}
		place.DeviceID = null.IntFromPtr(nil)
		places = append(places, place)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return places, err
	}

	return places, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdatePlace : update a place
func UpdatePlace(ctx context.Context, placeID int, zoneID int,
	placetype string, geo string, deviceID int) (PlaceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var place PlaceResponse

	place.PlaceID = -1

	if (zoneID == 0) && (placetype == "") && (geo == "") && (deviceID == 0) {
		return place, errors.New("invalid input fields (database/places.go")
	}

	// modify zoneID
	if zoneID != 0 {
		err := pool.QueryRowContext(ctx, `
			UPDATE places SET zone_id = $1 
			WHERE place_id = $2 RETURNING place_id
		`, zoneID, placeID).Scan(&place.PlaceID)

		if err == sql.ErrNoRows {
			log.Printf("no place with id %d\n", placeID)
			return place, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return place, err
		}
	}

	// modify type
	if placetype != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE places SET type = $1 
			WHERE place_id = $2 RETURNING place_id
		`, placetype, placeID).Scan(&place.PlaceID)

		if err == sql.ErrNoRows {
			log.Printf("no place with id %d\n", placeID)
			return place, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return place, err
		}
	}

	// modify geo
	if geo != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE places SET geo = $1 
			WHERE place_id = $2 RETURNING place_id
		`, geo, placeID).Scan(&place.PlaceID)

		if err == sql.ErrNoRows {
			log.Printf("no place with id %d\n", placeID)
			return place, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return place, err
		}
	}

	// modify deviceID
	if deviceID != 0 {
		err := pool.QueryRowContext(ctx, `
			UPDATE places SET device_id = $1 
			WHERE place_id = $2 RETURNING place_id
		`, deviceID, placeID).Scan(&place.PlaceID)

		if err == sql.ErrNoRows {
			log.Printf("no place with id %d\n", placeID)
			return place, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return place, err
		}
	}

	return place, nil
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/

// NewPlace : insert a new place
func NewPlace(ctx context.Context, zoneID int, placetype string,
	geo string, deviceID int) (PlaceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var place PlaceResponse

	place.PlaceID = -1

	err := pool.QueryRowContext(ctx,
		`INSERT INTO places
		(
			zone_id, 
			type,
			geo,
			device_id
		) VALUES
		(
			$1,
			$2,
			$3,
			$4
		) RETURNING place_id`, zoneID, placetype, geo, deviceID).Scan(&place.PlaceID)

	if err == sql.ErrNoRows {
		log.Printf("no place created\n")
		return place, err
	}

	if err != nil {
		log.Printf("query error: %v\n", err)
		return place, err
	}

	return place, nil
}

/********************************** CREATE **********************************/

/********************************** DELETE **********************************/

// DeletePlace : delete place
func DeletePlace(ctx context.Context, placeID int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// update the device id into places
	result, err := pool.ExecContext(ctx, `
		DELETE FROM places WHERE place_id = $1
	`, placeID)

	if err != nil {
		log.Printf("query error: %v\n", err)
		return false, err
	}

	return checkDeletion(result)
}

/********************************** DELETE **********************************/

/********************************** OPTIONS **********************************/

// CountPlace : count number of rows
func CountPlace(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var count int

	count = -1

	row := pool.QueryRowContext(ctx, "SELECT COUNT(*) FROM places")
	err := row.Scan(&count)
	if err != nil {
		log.Printf("query error: %v\n", err)
		return count, err
	}
	return count, nil
}

/********************************** OPTIONS **********************************/
