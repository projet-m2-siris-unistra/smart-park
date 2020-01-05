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
	database.Paging
	ZoneID int `json:"zone_id"`
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

type resultListPlace struct {
	Count int              `json:"count"`
	Data  []database.Place `json:"data"`
}

/********************************** GET **********************************/

func getPlace(ctx context.Context, request getPlaceRequest) (database.Place, error) {
	log.Println("handlers: handling getPlace")

	return database.GetPlace(ctx, request.PlaceID)
}

func getPlaces(ctx context.Context, request getPlacesRequest) (resultListPlace, error) {
	log.Println("handlers: handling getPlaces")

	var result resultListPlace
	var err error
	result.Count, err = database.CountPlace(ctx)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetPlaces(ctx, request.ZoneID, request.Paging)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getPlaceWithNoDevice(ctx context.Context) ([]database.Place, error) {
	log.Println("handlers: handling getPlaceWithNoDevice")

	return database.GetPlacesWithNoDevice(ctx)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

func updatePlace(ctx context.Context, request updatePlaceRequest) (database.PlaceResponse, error) {
	log.Println("handlers: handling updatePlace")

	return database.UpdatePlace(ctx, request.PlaceID, request.ZoneID,
		request.Type, request.Geography, request.DeviceID)
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/
func newPlace(ctx context.Context, request newPlaceRequest) (database.PlaceResponse, error) {
	log.Println("handlers: handling newPlace")

	return database.NewPlace(ctx, request.ZoneID,
		request.Type, request.Geography, request.DeviceID)
}

/********************************** CREATE **********************************/

/********************************** DELETE **********************************/

func deletePlace(ctx context.Context, request getPlaceRequest) (bool, error) {
	log.Println("handlers: handling deletePlace")

	return database.DeletePlace(ctx, request.PlaceID)
}

/********************************** DELETE **********************************/
