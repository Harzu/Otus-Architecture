package healthcheck

import (
	"net/http"
	"encoding/json"
	"otus-arch-hw/app/system/database/psql"
)

type HealthCheck struct {
	DBClient psql.Repository
}

func New(dbClient psql.Repository) HealthCheck {
	return HealthCheck{DBClient:dbClient}
}

type status struct {
	ServiceAvailable  bool
	DBAvailable       bool
}

func (hc *HealthCheck) Check(w http.ResponseWriter, req *http.Request) {
	data := status{}

	if !hc.DBClient.HealthCheck() {
		http.Error(w, "psql not connected", http.StatusInternalServerError)
		return
	}

	data.DBAvailable = true
	data.ServiceAvailable = true

	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}