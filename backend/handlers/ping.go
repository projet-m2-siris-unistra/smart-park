package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

func ping(ctx context.Context) (string, error) {
	log.Println("handlers: handling ping")

	ret, err := database.Ping(ctx)
	if err != nil {
		return "", err
	}

	return ret, nil
}
