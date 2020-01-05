package utils

import (
	"strconv"
)

// PageInfo holds informations about the current page
type PageInfo struct {
	Total   int    `json:"total"`
	HasNext bool   `json:"has_next"`
	HasPrev bool   `json:"has_prev"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

// GeneratePageInfo returns a PageInfo object given the total, offset and limit
func GeneratePageInfo(total, offset, limit int) PageInfo {
	hasPrev := offset > 0
	hasNext := (limit + offset) < total
	return PageInfo{
		Total:   total,
		HasNext: hasNext,
		HasPrev: hasPrev,
		Limit:   limit,
		Offset:  offset,
	}
}

// ParseOffsetLimit returns a offset & limit given the request args
func ParseOffsetLimit(vars map[string]string) (int, int) {
	var offset, limit int
	offsetStr, ok := vars["offset"]
	if !ok {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(offsetStr)
	}

	if offset < 0 {
		offset = 10
	}

	limitStr, ok := vars["limit"]
	if !ok {
		limit = 10
	} else {
		limit, _ = strconv.Atoi(limitStr)
	}

	if limit < 0 || limit > 100 {
		limit = 10
	}

	return offset, limit
}
