package bus

import (
	"context"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
	"github.com/projet-m2-siris-unistra/smart-park/backend/handlers"
)

// ListZones returns the list of zones of a tenant
func ListZones(ctx context.Context, tenantID int, offset, limit int) (*handlers.ZoneList, error) {
	req := handlers.ListZonesRequest{
		TenantID: tenantID,
		Paging: database.Paging{
			Limit:  limit,
			Offset: offset,
		},
		ZoneOptions: database.ZoneOptions{
			WithPlaces: true,
		},
	}
	resp := handlers.ZoneList{}
	err := jsonConn.RequestWithContext(ctx, "zones.list", &req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZone fetches informations about a zone
func GetZone(ctx context.Context, tenantID int, zoneID int) (*database.Zone, error) {
	req := handlers.GetZoneRequest{
		ZoneID: zoneID,
		ZoneOptions: database.ZoneOptions{
			WithPlaces: true,
		},
	}
	resp := database.Zone{}
	err := jsonConn.RequestWithContext(ctx, "zones.get", &req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
