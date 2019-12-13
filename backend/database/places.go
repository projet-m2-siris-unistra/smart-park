package database

import (
	"context"
	"errors"
	"log"
	"time"
	"database/sql"

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
func GetPlaces(ctx context.Context, zoneID int, limite int, offset int) ([]Place, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var places []Place
	var place Place
	var i int

	limite, offset = CheckArgPlace(limite, offset)

	i = 0
	
	rows, err := pool.QueryContext(ctx,
		`SELECT place_id, zone_id, type, geo, place_id, created_at, updated_at
		FROM places WHERE zone_id = $1 LIMIT $2 OFFSET $3`, zoneID, limite, offset)

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
			log.Printf("no place with id %d\n", deviceID)
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
			log.Printf("no place with id %d\n", deviceID)
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
			log.Printf("no place with id %d\n", deviceID)
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
			log.Printf("no place with id %d\n", deviceID)
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


/********************************** OPTIONS **********************************/

// CheckArgPlace : check limit and offset arguments
func CheckArgPlace(limite int, offset int) (int, int) {

	if limite == 0 {
		limite = 20
	}

	if offset == 0 {
		offset = 0
	}

	return limite, offset
}

/********************************** OPTIONS **********************************/
