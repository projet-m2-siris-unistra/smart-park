package zones

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/guregu/null.v3"

	"github.com/projet-m2-siris-unistra/smart-park/api/bus"
	"github.com/projet-m2-siris-unistra/smart-park/api/utils"
	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

const tenantID = 1

type point struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type zone struct {
	ZoneID    int                  `json:"zone_id"`
	TenantID  int                  `json:"tenant_id"`
	Name      string               `json:"name"`
	Type      database.ZoneType    `json:"type"`
	Color     null.String          `json:"color"`
	Geography []point              `json:"geo"`
	Places    *database.ZonePlaces `json:"places"`
	database.Timestamps
}

func parseGeography(geo string) []point {
	re := regexp.MustCompile(`\[\s*(\d*\.\d*)\s*,\s*(\d*.\d*)\s*\]`)
	matches := re.FindAllStringSubmatch(geo, -1)
	ret := make([]point, len(matches))
	for i, match := range matches {
		lat, _ := strconv.ParseFloat(match[1], 64)
		long, _ := strconv.ParseFloat(match[2], 64)
		ret[i].Lat = lat
		ret[i].Long = long
	}
	return ret
}

func mapZone(z *database.Zone) zone {
	geo := []point{}
	if !z.Geography.IsZero() {
		geo = parseGeography(*z.Geography.Ptr())
	}

	return zone{
		ZoneID:     z.ZoneID,
		TenantID:   z.TenantID,
		Name:       z.Name,
		Type:       z.Type,
		Color:      z.Color,
		Geography:  geo,
		Places:     z.Places,
		Timestamps: z.Timestamps,
	}
}

type zoneList struct {
	Info  utils.PageInfo `json:"page"`
	Zones []zone         `json:"zones"`
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	offset, limit := utils.ParseOffsetLimit(vars)
	list, err := bus.ListZones(ctx, tenantID, offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respList := zoneList{
		Info:  utils.GeneratePageInfo(list.Count, offset, limit),
		Zones: []zone{},
	}

	for _, z := range list.Data {
		respList.Zones = append(respList.Zones, mapZone(&z))
	}

	resp, err := json.Marshal(respList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	zone, err := bus.GetZone(ctx, tenantID, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(mapZone(zone))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// Register adds the zones routes to the giver router
func Register(router *mux.Router) {
	s := router.PathPrefix("/zones").Subrouter()
	s.Path("").
		Queries(
			"offset", "{offset:[0-9]+}",
			"limit", "{limit:[0-9]+}",
		).
		Methods("GET").
		HandlerFunc(index)

	s.Path("").
		Methods("GET").
		HandlerFunc(index)

	s.Path("/{id:[0-9]+}").
		Methods("GET").
		HandlerFunc(get)
}
