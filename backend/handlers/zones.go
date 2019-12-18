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
	Limite int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
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

type resultListZone struct {
	Count int `json:"count"`
	Data []database.Zone `json:"data"`
}

/********************************** GET **********************************/

func getZone(ctx context.Context, request getZoneRequest) (database.Zone, error) {
	log.Println("handlers: handling getZone")

	return database.GetZone(ctx, request.ZoneID)
}

func getZones(ctx context.Context, request getZonesRequest) (resultListZone, error) {
	log.Println("handlers: handling getZones of a tenant")

	var result resultListZone
	var err error 
	result.Count, err = database.CountZone(ctx, request.TenantID)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetZones(ctx, request.TenantID, request.Limite, request.Offset)
	if err != nil {
		return result, err
	}
	return result, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

func updateZone(ctx context.Context, request updateZoneRequest) (database.ZoneResponse, error) {
	log.Println("handlers: handling updateZone")

	return database.UpdateZone(ctx, request.ZoneID, request.TenantID,
		request.Name, request.Type, request.Color, request.Geography)
}

/********************************** UPDATE **********************************/

/********************************** CREATE **********************************/
func newZone(ctx context.Context, request newZoneRequest) (database.ZoneResponse, error) {
	log.Println("handlers: handling newZone")

	return database.NewZone(ctx, request.TenantID,
		request.Name, request.Type, request.Color, request.Geography)
}

/********************************** CREATE **********************************/

/********************************** DELETE **********************************/
func deleteZone(ctx context.Context, request getZoneRequest) (database.ZoneResponse, error) {
	log.Println("handlers: handling deleteZone")

	return database.DeleteZone(ctx, request.ZoneID)
}

/********************************** DELETE **********************************/