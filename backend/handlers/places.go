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
}

func getPlace(ctx context.Context, request getPlaceRequest) (database.Place, error) {
	log.Println("handlers: handling getPlace")

	return database.GetPlace(ctx, request.PlaceID)
}

func getPlaces(ctx context.Context, request getPlacesRequest) ([]database.Place, error) {
	log.Println("handlers: handling getPlaces")

	return database.GetPlaces(ctx, request.ZoneID)
}
