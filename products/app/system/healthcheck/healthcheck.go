package healthcheck

import (
	"encoding/json"
	"net/http"
	"otus-auth/app/system/database/psql"
	"otus-auth/app/system/database/redis"
)

type HealthCheck struct {
	DBClient    psql.Repository
	RedisClient redis.Repository
}

func New(dbClient psql.Repository, redisClient redis.Repository) HealthCheck {
	return HealthCheck{DBClient: dbClient, RedisClient: redisClient}
}

type status struct {
	DBAvailable    bool
	RedisAvailable bool
}

func (hc *HealthCheck) Check(w http.ResponseWriter, req *http.Request) {
	data := status{}

	if hc.DBClient.HealthCheck() {
		data.DBAvailable = true
	}

	if hc.RedisClient.HealthCheck() {
		data.RedisAvailable = true
	}

	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
