package tenants

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/projet-m2-siris-unistra/smart-park/api/bus"
	"github.com/projet-m2-siris-unistra/smart-park/api/utils"
	"github.com/projet-m2-siris-unistra/smart-park/api/v1/zones"
	"github.com/projet-m2-siris-unistra/smart-park/backend/database"
)

type tenantList struct {
	Info    utils.PageInfo    `json:"page"`
	Tenants []database.Tenant `json:"tenants"`
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	vars := mux.Vars(r)
	offset, limit := utils.ParseOffsetLimit(vars)
	list, err := bus.ListTenants(ctx, offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respList := tenantList{
		Info:    utils.GeneratePageInfo(list.Count, offset, limit),
		Tenants: []database.Tenant{},
	}

	for _, item := range list.Data {
		respList.Tenants = append(respList.Tenants, item)
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
	id, _ := strconv.Atoi(vars["tenant_id"])
	tenant, err := bus.GetTenant(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(tenant)
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
	root := router.PathPrefix("/tenants").Subrouter()
	root.Path("").
		Queries(
			"offset", "{offset:[0-9]+}",
			"limit", "{limit:[0-9]+}",
		).
		Methods("GET").
		HandlerFunc(index)

	root.Path("").
		Methods("GET").
		HandlerFunc(index)

	single := root.PathPrefix("/{tenant_id:[0-9]+}").Subrouter()
	single.Path("").
		Methods("GET").
		HandlerFunc(get)
	zones.Register(single)
}
