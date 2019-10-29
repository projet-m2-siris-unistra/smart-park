package handlers

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

func ping(m *nats.Msg) {
	log.Println("handlers: handling ping")
	ret, err := database.Ping(context.TODO())
	if err != nil {
		return
	}

	m.Respond([]byte(ret))
}
