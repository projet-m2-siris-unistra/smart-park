package database

import (
	"fmt"
)

const maxLimit = 100
const defaultLimit = 10

// Paging holds paging (limit & offset) informations
type Paging struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// Normalize checks that values are in bound
func (p Paging) Normalize() Paging {
	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
	if p.Limit > maxLimit {
		p.Limit = maxLimit
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
	return p
}

func (p Paging) buildQuery() string {
	return fmt.Sprintf("LIMIT %d OFFSET %d", p.Limit, p.Offset)
}
