package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type GetTenantRequest struct {
	TenantID int `json:"tenant_id"`
}

type ListTenantsRequest struct {
	database.Paging
}

type updateTenantsRequest struct {
	TenantID  int    `json:"tenant_id"`
	Name      string `json:"name,omitempty"`
	Geography string `json:"geo,omitempty"`
}

type TenantList struct {
	Count int               `json:"count"`
	Data  []database.Tenant `json:"data"`
}

/********************************** GET **********************************/
func getTenant(ctx context.Context, request GetTenantRequest) (database.Tenant, error) {
	log.Println("handlers: handling getTenant")

	return database.GetTenant(ctx, request.TenantID)
}

func getTenants(ctx context.Context, request ListTenantsRequest) (TenantList, error) {
	log.Println("handlers: handling getTenants")

	var result TenantList
	var err error
	result.Count, err = database.CountTenant(ctx)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetTenants(ctx, request.Paging)
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
