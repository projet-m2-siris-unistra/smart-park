package handlers

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

func ping(m *nats.Msg) {
	ctx := context.TODO()
	log.Println("handlers: handling ping")

	ret, err := database.Ping(ctx)
	if err != nil {
		return
	}

	m.Respond([]byte(ret))
}
