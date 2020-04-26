package services

import (
	"otus-arch-hw/app/services/metrics"
)

type Services struct {
	Metrics metrics.Client
}

func New() *Services {
	return &Services{
		Metrics: metrics.New(),
	}
}