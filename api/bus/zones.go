package bus

import (
	"context"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
	"github.com/projet-m2-siris-unistra/smart-park/backend/handlers"
)

// ListZones returns the list of zones of a handler
func ListZones(ctx context.Context, tenantID int, offset, limit int) (*handlers.ZoneList, error) {
	req := handlers.ListZonesRequest{
		TenantID: tenantID,
		Limit:    limit,
		Offset:   offset,
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
	req := handlers.GetZoneRequest{ZoneID: zoneID}
	resp := database.Zone{}
	err := jsonConn.RequestWithContext(ctx, "zones.get", &req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
