package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getPlaceRequest struct {
	PlaceID int `json:"place_id"`
}

type getPlacesRequest struct {
	ZoneID int `json:"zone_id"`
	Limite int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type updatePlaceRequest struct {
	PlaceID   int    `json:"place_id"`
	ZoneID    int    `json:"zone_id,omitempty"`
	Type      string `json:"type,omitempty"`
	Geography string `json:"geo,omitempty"`
	DeviceID  int    `json:"device_id,omitempty"`
}

type newPlaceRequest struct {
	ZoneID    int    `json:"zone_id"`
	Type      string `json:"type"`
	Geography string `json:"geo"`
	DeviceID  int    `json:"device_id"`
}

/********************************** GET **********************************/

func getPlace(ctx context.Context, request getPlaceRequest) (database.Place, error) {
	log.Println("handlers: handling getPlace")

	return database.GetPlace(ctx, request.PlaceID)
}

func getPlaces(ctx context.Context, request getPlacesRequest) ([]database.Place, error) {
	log.Println("handlers: handling getPlaces")

	return database.GetPlaces(ctx, request.ZoneID, request.Limite, request.Offset)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

func updatePlace(ctx context.Context, request updatePlaceRequest) error {
	log.Println("handlers: handling updatePlace")

	err := database.UpdatePlace(ctx, request.PlaceID, request.ZoneID,
		request.Type, request.Geography, request.DeviceID)
	return err
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/
func newPlace(ctx context.Context, request newPlaceRequest) error {
	log.Println("handlers: handling newPlace")

	err := database.NewPlace(ctx, request.ZoneID,
		request.Type, request.Geography, request.DeviceID)
	return err
}

/********************************** CREATE **********************************/
