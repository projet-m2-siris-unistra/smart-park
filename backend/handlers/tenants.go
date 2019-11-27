package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getTenantRequest struct {
	TenantID int `json:"tenant_id"`
}

type updateGeoTenantsRequest struct {
	TenantID  int    `json:"tenant_id"`
	Geography string `json:"geo"`
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

func updateGeoTenants(ctx context.Context, request updateGeoTenantsRequest) error {
	log.Println("handlers: handling updateGeoTenants")

	err := database.UpdateGeoTenants(ctx, request.TenantID, request.Geography)
	return err
}

/********************************** UPDATE **********************************/
