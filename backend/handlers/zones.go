package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getZoneRequest struct {
	ZoneID int `json:"zone_id"`
}

// get all Zone of a tenant
type getZonesRequest struct {
	TenantID int `json:"tenant_id"`
}

func getZone(ctx context.Context, request getZoneRequest) (database.Zone, error) {
	log.Println("handlers: handling getZone")

	return database.GetZone(ctx, request.ZoneID)
}

func getZones(ctx context.Context, request getZonesRequest) ([]database.Zone, error) {
	log.Println("handlers: handling getZones of a tenant")

	return database.GetZones(ctx, request.TenantID)
}
