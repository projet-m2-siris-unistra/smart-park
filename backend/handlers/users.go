package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getUserRequest struct {
	UserID int `json:"user_id"`
}

type getUsersRequest struct {
	Limite int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type updateUserRequest struct {
	UserID   int    `json:"user_id"`
	TenantID int    `json:"tenant_id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type resultListUser struct {
	Count int `json:"count"`
	Data []database.User `json:"data"`
}

/********************************** GET **********************************/

func getUser(ctx context.Context, request getUserRequest) (database.User, error) {
	log.Println("handlers: handling getUser")

	return database.GetUser(ctx, request.UserID)
}

func getUsers(ctx context.Context, request getUsersRequest) (resultListUser, error) {
	log.Println("handlers: handling getUsers")

	var result resultListUser
	var err error 
	result.Count, err = database.CountUser(ctx)
	if err != nil {
		return result, err
	}
	result.Data, err = database.GetUsers(ctx, request.Limite, request.Offset)
	if err != nil {
		return result, err
	}
	return result, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/
func updateUser(ctx context.Context, request updateUserRequest) (database.UserResponse, error)  {
	log.Println("handlers: handling updateUser")

	return database.UpdateUser(ctx, request.UserID, request.TenantID, request.Username,
		request.Password, request.Email)
}

/********************************** UPDATE **********************************/
