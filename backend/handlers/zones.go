package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getZoneRequest struct {
	ZoneID int `json:"zone_id"`
}

func getZone(ctx context.Context, request getZoneRequest) (database.Zone, error) {
	log.Println("handlers: handling getZone")

	return database.GetZone(ctx, request.ZoneID)
}
