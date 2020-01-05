package bus

import (
	"context"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
	"github.com/projet-m2-siris-unistra/smart-park/backend/handlers"
)

// ListTenants returns the list of tenants
func ListTenants(ctx context.Context, offset, limit int) (*handlers.TenantList, error) {
	req := handlers.ListTenantsRequest{
		Paging: database.Paging{
			Limit:  limit,
			Offset: offset,
		},
	}
	resp := handlers.TenantList{}
	err := jsonConn.RequestWithContext(ctx, "tenants.list", &req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTenant fetches informations about a tenant
func GetTenant(ctx context.Context, tenantID int) (*database.Tenant, error) {
	req := handlers.GetTenantRequest{
		TenantID: tenantID,
	}
	resp := database.Tenant{}
	err := jsonConn.RequestWithContext(ctx, "tenants.get", &req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
