package database

import (
	"context"
	"time"
)

// Tenant : owner of one park
type Tenant struct {
	TenantID  int       `json:"id"`
	Name      string    `json:"name"`
	Geography string    `json:"geo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetTenant fetches the tenant by its ID
func GetTenant(ctx context.Context, tenantID int) (Tenant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var tenant Tenant

	err := pool.QueryRowContext(ctx, `
		SELECT tenant_id, name, geo, created_at, updated_at
		FROM tenants 
		WHERE tenant_id = $1
	`, tenantID).
		Scan(&tenant.TenantID, &tenant.Name, &tenant.Geography,
			&tenant.CreatedAt, &tenant.UpdatedAt)

	if err != nil {
		return tenant, err
	}

	return tenant, nil
}
