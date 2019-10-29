package database

import (
	"context"
	"time"
)

// Ping the database
func Ping(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var ret string
	err := pool.QueryRowContext(ctx, "SELECT 'PONG';").Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}
