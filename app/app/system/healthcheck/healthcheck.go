package healthcheck

import (
	"net/http"
	"encoding/json"
)

type HealthCheck struct {
}

func New() HealthCheck {
	return HealthCheck{}
}

type status struct {
	ServiceAvailable  bool
}

func (hc *HealthCheck) Check(w http.ResponseWriter, req *http.Request) {
	data := status{}
	data.ServiceAvailable = true
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}