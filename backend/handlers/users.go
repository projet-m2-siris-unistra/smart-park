package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getUserRequest struct {
	UserID int `json:"user_id"`
}

type updateUserRequest struct {
	UserID   int    `json:"user_id"`
	TenantID int    `json:"tenant_id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

/********************************** GET **********************************/

func getUser(ctx context.Context, request getUserRequest) (database.User, error) {
	log.Println("handlers: handling getUser")

	return database.GetUser(ctx, request.UserID)
}

func getUsers(ctx context.Context, request getUserRequest) ([]database.User, error) {
	log.Println("handlers: handling getUsers")

	return database.GetUsers(ctx)
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/
func updateUser(ctx context.Context, request updateUserRequest) error {
	log.Println("handlers: handling updateUser")

	err := database.UpdateUser(ctx, request.UserID, request.TenantID, request.Username,
		request.Password, request.Email)
	return err
}

/********************************** UPDATE **********************************/
