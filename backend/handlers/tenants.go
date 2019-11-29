package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getTenantRequest struct {
	TenantID int `json:"tenant_id"`
}

type updateTenantsRequest struct {
	TenantID  int    `json:"tenant_id"`
	Name      string `json:"name,omitempty"`
	Geography string `json:"geo,omitempty"`
}

/********************************** GET **********************************/
func getTenant(ctx context.Context, request getTenantRequest) (database.Tenant, error) {
	log.Println("handlers: handling getTenant")

	return database.GetTenant(ctx, request.TenantID)
}

func getTenants(ctx context.Context, request getTenantRequest) ([]database.Tenant, error) {
	log.Println("handlers: handling getTenants")

	return database.GetTenants(ctx)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

func updateTenants(ctx context.Context, request updateTenantsRequest) error {
	log.Println("handlers: handling updateGeoTenants")

	err := database.UpdateTenants(ctx, request.TenantID, request.Name, request.Geography)
	return err
}

/********************************** UPDATE **********************************/
