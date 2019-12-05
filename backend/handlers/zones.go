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

type updateZoneRequest struct {
	ZoneID    int    `json:"zone_id"`
	TenantID  int    `json:"tenant_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	Color     string `json:"color,omitempty"`
	Geography string `json:"geo,omitempty"`
}

type newZoneRequest struct {
	TenantID  int    `json:"tenant_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Color     string `json:"color"`
	Geography string `json:"geo"`
}

/********************************** GET **********************************/

func getZone(ctx context.Context, request getZoneRequest) (database.Zone, error) {
	log.Println("handlers: handling getZone")

	return database.GetZone(ctx, request.ZoneID)
}

func getZones(ctx context.Context, request getZonesRequest) ([]database.Zone, error) {
	log.Println("handlers: handling getZones of a tenant")

	return database.GetZones(ctx, request.TenantID)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

func updateZone(ctx context.Context, request updateZoneRequest) error {
	log.Println("handlers: handling updateZone")

	err := database.UpdateZone(ctx, request.ZoneID, request.TenantID,
		request.Name, request.Type, request.Color, request.Geography)
	return err
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/
func newZone(ctx context.Context, request newZoneRequest) error {
	log.Println("handlers: handling newZone")

	err := database.NewZone(ctx, request.TenantID,
		request.Name, request.Type, request.Color, request.Geography)
	return err
}

/********************************** CREATE **********************************/