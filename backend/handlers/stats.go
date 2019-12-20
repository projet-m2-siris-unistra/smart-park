package handlers

import (
	"context"
	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
	"log"
)

/*type statsParameters struct {
	PlaceID int `json:"device_id,omitempty"`
}*/

func getStatsDevice(ctx context.Context) ([]int, []int, []int, error) {
	log.Println("handlers: handling getStatsDevice")

	//var devices []int
	//var stateDevices []int
	//var zones []int

	return database.GetStatsDevice(ctx)
}
