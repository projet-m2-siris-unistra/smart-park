package handlers

import (
	"context"
	"log"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type newFakeDataRequest struct {
	Tenants int `json:"tenants,omitempty"`
	Zones   int `json:"zones,omitempty"`
	Devices int `json:"devices,omitempty"`
	Places  int `json:"places,omitempty"`
	Users   int `json:"users,omitempty"`
}

func createFakeData(ctx context.Context, request newFakeDataRequest) error {
	log.Println("handlers: handling createFakeData")

	return database.Faker(ctx, request.Tenants, request.Zones, request.Devices,
		request.Places, request.Users)
}
