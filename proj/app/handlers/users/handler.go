package users

import (
	userRepo "otus-arch-hw/app/repositories/user"
	"otus-arch-hw/app/services/metrics"
	"time"
)

type UserHandler struct {
	metrics    metrics.Client
	Repository userRepo.Repository
}

func NewHandler(repo userRepo.Repository, metrics metrics.Client) UserHandler {
	return UserHandler{
		metrics:    metrics,
		Repository: repo,
	}
}

func (u *UserHandler) afterRequest(handler string, method string, statusCode int, startedAt time.Time) {
	u.metrics.LatencyMetric.CountLatency(handler, method, startedAt)
	u.metrics.HttpCountMetric.Count(handler, method, statusCode)
}
