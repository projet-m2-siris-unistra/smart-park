package database

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v3"
)

// Tenant : owner of one park
type Tenant struct {
	TenantID  int         `json:"id"`
	Name      string      `json:"name"`
	Geography null.String `json:"geo"`
	Timestamps
}

// GetTenant : fetches the tenant by its ID
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

// GetTenants : get all the tenant
func GetTenants(ctx context.Context) ([]Tenant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var tenants []Tenant
	var tenant Tenant
	var i int

	i = 0
	rows, err := pool.QueryContext(ctx,
		`SELECT tenant_id, name, geo, created_at, updated_at FROM tenants`)

	if err != nil {
		return tenants, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tenant.TenantID, &tenant.Name, &tenant.Geography,
			&tenant.CreatedAt, &tenant.UpdatedAt)
		if err != nil {
			return tenants, err
		}
		tenants = append(tenants, tenant)
		i = i + 1
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return tenants, err
	}

	return tenants, nil
}
