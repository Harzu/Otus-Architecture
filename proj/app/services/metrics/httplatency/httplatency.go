package httplatency

import (
	"time"
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	labelHandlerName = "handler_name"
	labelMethodName  = "method_name"
)

type HttpLatency struct {
	*sync.Mutex
	metric *prometheus.HistogramVec
	values map[string]metricValue
}

func NewHttpLatency() *HttpLatency {
	return &HttpLatency{
		Mutex:  &sync.Mutex{},
		values: make(map[string]metricValue),
		metric: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:        "app_request_duration_time",
			Help:        "latency for http handlers by buckets",
		}, []string{labelHandlerName, labelMethodName}),
	}
}

func (h *HttpLatency) Register(registry *prometheus.Registry) {
	registry.Register(h.metric)
}

func (h *HttpLatency) SetToPrometheus() {
	h.Lock()
	defer h.Unlock()

	for _, val := range h.values {
		h.metric.With(prometheus.Labels{
			labelMethodName: val.methodName,
			labelHandlerName: val.handlerName,
		}).Observe(float64(val.value))
	}
}

type metricValue struct {
	handlerName string
	methodName  string
	value       int64
}

func (h *HttpLatency) CountLatency(handler, method string, started time.Time) {
	h.Lock()
	defer h.Unlock()

	duration := time.Since(started).Milliseconds()
	key := fmt.Sprintf("%s_%s", handler, method)
	h.values[key] = metricValue{
		handlerName: handler,
		methodName:  method,
		value:       duration,
	}
}