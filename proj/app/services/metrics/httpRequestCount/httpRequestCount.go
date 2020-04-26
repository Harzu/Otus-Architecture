package httpRequestCount

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"fmt"
	"strconv"
)

const (
	labelHandlerName = "handler_name"
	labelMethodName  = "method_name"
	labelStatusCode  = "status_code"
)

type HttpCount struct {
	*sync.Mutex
	metric *prometheus.GaugeVec
	values map[string]metricValue
}

func NewHttpCount() *HttpCount {
	return &HttpCount{
		Mutex:  &sync.Mutex{},
		values: make(map[string]metricValue),
		metric: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:        "app_request_count",
			Help:        "count for http handlers",
		}, []string{labelHandlerName, labelMethodName, labelStatusCode}),
	}
}

func (h *HttpCount) Register(registry *prometheus.Registry) {
	registry.Register(h.metric)
}

func (h *HttpCount) SetToPrometheus() {
	h.Lock()
	defer h.Unlock()

	h.metric.Reset()
	for _, val := range h.values {
		h.metric.With(prometheus.Labels{
			labelMethodName: val.methodName,
			labelHandlerName: val.handlerName,
			labelStatusCode: val.statusCode,
		}).Set(float64(val.value))
	}

	h.values = make(map[string]metricValue)
}

type metricValue struct {
	handlerName string
	methodName  string
	statusCode  string
	value       int64
}

func (h *HttpCount) Count(handler, method string, status int) {
	h.Lock()
	defer h.Unlock()

	key := fmt.Sprintf("%s_%s_%d", handler, method, status)
	target, ok := h.values[key]
	if !ok {
		target = metricValue{
			handlerName: handler,
			methodName:  method,
			statusCode:  strconv.Itoa(status),
			value:       1,
		}
	} else {
		target.value++
	}

	h.values[key] = target
}