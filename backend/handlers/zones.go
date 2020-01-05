package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

// GetZoneRequest holds the parameters of a zones.get request
type GetZoneRequest struct {
	ZoneID         int  `json:"zone_id"`
	WithPlacesInfo bool `json:"with_places_info"`
}

// ListZonesRequest holds the parameters of a zones.list request
type ListZonesRequest struct {
	TenantID int `json:"tenant_id"`
	Limit    int `json:"limit,omitempty"`
	Offset   int `json:"offset,omitempty"`
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

// ZoneList holds the result of a zones.list call
type ZoneList struct {
	Count int             `json:"count"`
	Data  []database.Zone `json:"data"`
}

/********************************** GET **********************************/

func getZone(ctx context.Context, request GetZoneRequest) (database.Zone, error) {
	log.Println("handlers: handling getZone")

	return database.GetZone(ctx, request.ZoneID)
}

func getZones(ctx context.Context, request ListZonesRequest) (ZoneList, error) {
	log.Println("handlers: handling getZones of a tenant")

	var result ZoneList
	var err error
	result.Count, err = database.CountZone(ctx, request.TenantID)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetZones(ctx, request.TenantID, request.Limit, request.Offset)
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
func deleteZone(ctx context.Context, request GetZoneRequest) (database.ZoneResponse, error) {
	log.Println("handlers: handling deleteZone")

	return database.DeleteZone(ctx, request.ZoneID)
}

/********************************** DELETE **********************************/
