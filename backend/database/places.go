package database

import (
	"context"
	"errors"
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
	DeviceID  int         `json:"device_id"`
	Timestamps
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

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdatePlace : update a place
func UpdatePlace(ctx context.Context, placeID int, zoneID int,
	placetype string, geo string, deviceID int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if (zoneID == 0) && (placetype == "") && (geo == "") && (deviceID == 0) {
		return errors.New("invalid input fields (database/places.go")
	}

	// modify zoneID
	if zoneID != 0 {
		result, err := pool.ExecContext(ctx, `
			UPDATE places SET zone_id = $1 
			WHERE place_id = $2
		`, zoneID, placeID)

		if err != nil {
			return errors.New("error update place zone_id")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : place zone_id - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify type
	if placetype != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE places SET type = $1 
			WHERE place_id = $2
		`, placetype, placeID)

		if err != nil {
			return errors.New("error update place type")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : place type - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify geo
	if geo != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE places SET geo = $1 
			WHERE place_id = $2
		`, geo, placeID)

		if err != nil {
			return errors.New("error update place geo")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : place geo - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify deviceID
	if deviceID != 0 {
		result, err := pool.ExecContext(ctx, `
			UPDATE places SET device_id = $1 
			WHERE place_id = $2
		`, deviceID, placeID)

		if err != nil {
			return errors.New("error update place device_id")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : place device_id - rows affected")
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

// NewPlace : insert a new place
func NewPlace(ctx context.Context, zoneID int, placetype string,
	geo string, deviceID int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := pool.ExecContext(ctx,
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
		)`, zoneID, placetype, geo, deviceID)

	if err != nil {
		return errors.New("error new place")
	}

	return nil
}

/********************************** CREATE **********************************/
