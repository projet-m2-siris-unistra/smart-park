package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type getTenantRequest struct {
	TenantID int `json:"tenant_id"`
}

func getTenant(m *nats.Msg) {
	ctx := context.TODO()
	log.Println("handlers: handling getTenant")

	var request getTenantRequest
	err := json.Unmarshal(m.Data, &request)
	if err != nil {
		log.Println(err)
		return
	}

	tenant, err := database.GetTenant(ctx, request.TenantID)
	if err != nil {
		log.Println(err)
		return
	}

	payload, err := json.Marshal(tenant)
	if err != nil {
		log.Println(err)
		return
	}

	m.Respond(payload)
}
