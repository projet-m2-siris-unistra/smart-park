package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getTenantRequest struct {
	TenantID int `json:"tenant_id"`
}

type getTenantsRequest struct {
	Limite int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type updateTenantsRequest struct {
	TenantID  int    `json:"tenant_id"`
	Name      string `json:"name,omitempty"`
	Geography string `json:"geo,omitempty"`
}

type resultListTenant struct {
	Count int `json:"count"`
	Data []database.Tenant `json:"data"`
}

/********************************** GET **********************************/
func getTenant(ctx context.Context, request getTenantRequest) (database.Tenant, error) {
	log.Println("handlers: handling getTenant")

	return database.GetTenant(ctx, request.TenantID)
}

func getTenants(ctx context.Context, request getTenantsRequest) (resultListTenant, error) {
	log.Println("handlers: handling getTenants")

	var result resultListTenant
	var err error 
	result.Count, err = database.CountTenant(ctx)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetTenants(ctx, request.Limite, request.Offset)
	if err != nil {
		return result, err
	}
	return result, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

func updateTenants(ctx context.Context, request updateTenantsRequest) (database.TenantResponse, error) {
	log.Println("handlers: handling updateGeoTenants")

	return database.UpdateTenants(ctx, request.TenantID, request.Name, request.Geography)
}

/********************************** UPDATE **********************************/
