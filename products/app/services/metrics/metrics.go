package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var metricsRegistry = prometheus.NewRegistry()

type NopMetric struct{}

func NewNopMetric() *NopMetric {
	return &NopMetric{}
}

func (NopMetric) MustRegister(...prometheus.Collector) {
}

type PromMetric interface {
	// SetToPrometheus is expected to bulk-update specific metrics.
	SetToPrometheus()
}

type Client struct{}

func New() Client {
	c := Client{}

	return c
}

func (c Client) Metrics() []PromMetric {
	return []PromMetric{}
}

func NewMetricsHandler(metrics []PromMetric) http.Handler {
	return newHandler(metricsRegistry, metrics)
}

func newHandler(registry *prometheus.Registry, metrics []PromMetric) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		for _, m := range metrics {
			m.SetToPrometheus()
		}
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(rsp, req)
	})
}

func init() {
	metricsRegistry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	metricsRegistry.MustRegister(prometheus.NewGoCollector())
}
