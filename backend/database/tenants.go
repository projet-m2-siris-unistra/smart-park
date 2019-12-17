package database

import (
	"context"
	"errors"
	"log"
	"time"
	"database/sql"

	"gopkg.in/guregu/null.v3"
)

// Tenant : owner of one park
type Tenant struct {
	TenantID  int         `json:"id"`
	Name      string      `json:"name"`
	Geography null.String `json:"geo"`
	Timestamps
}

// TenantResponse returns the id of the updated / created object 
type TenantResponse struct {
	TenantID int `json:"tenant_id"`
}

/********************************** GET **********************************/

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
func GetTenants(ctx context.Context, limite int, offset int) ([]Tenant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var tenants []Tenant
	var tenant Tenant
	var i int
 
	limite, offset = CheckArgTenant(limite, offset)

	i = 0

	rows, err := pool.QueryContext(ctx,
		`SELECT DISTINCT tenant_id, name, geo, created_at, updated_at 
		FROM tenants ORDER BY tenant_id LIMIT $1 OFFSET $2`, limite, offset)

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

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdateTenants : update a tenant
func UpdateTenants(ctx context.Context, tenantID int, name string, geo string) (TenantResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var tenant TenantResponse

	tenant.TenantID = -1

	if (name == "") && (geo == "") {
		return tenant, errors.New("invalid input fields (database/tenants.go)")
	}

	if geo != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE tenants SET geo = $1 
			WHERE tenant_id = $2 RETURNING tenant_id
		`, geo, tenantID).Scan(&tenant.TenantID)

		if err == sql.ErrNoRows {
			log.Printf("no device with id %d\n", tenantID)
			return tenant, err
		}
	
		if err != nil {
			log.Printf("query error: %v\n", err)
			return tenant, err
		}
	}
	
	if name != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE tenants SET name = $1 
			WHERE tenant_id = $2 RETURNING tenant_id
		`, name, tenantID).Scan(&tenant.TenantID)

		if err == sql.ErrNoRows {
			log.Printf("no device with id %d\n", tenantID)
			return tenant, err
		}
	
		if err != nil {
			log.Printf("query error: %v\n", err)
			return tenant, err
		}
	}

	return tenant, nil
}

/********************************** UPDATE **********************************/

/********************************** OPTIONS **********************************/

// CountTenant : count number of rows
func CountTenant(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	var count int

	count = -1

	row := pool.QueryRow("SELECT COUNT(*) FROM tenants")
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// CheckArgTenant : check limit and offset arguments
func CheckArgTenant(limite int, offset int) (int, int) {

	if limite == 0 {
		limite = 20
	}

	if offset == 0 {
		offset = 0
	}

	return limite, offset
}

/********************************** OPTIONS **********************************/
