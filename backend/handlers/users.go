package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getUserRequest struct {
	UserID int `json:"user_id"`
}

func getUser(ctx context.Context, request getUserRequest) (database.User, error) {
	log.Println("handlers: handling getUser")

	return database.GetUser(ctx, request.UserID)
}

func getUsers(ctx context.Context, request getUserRequest) ([]database.User, error) {
	log.Println("handlers: handling getUsers")

	return database.GetUsers(ctx)
}
