package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsHttp struct {
	count prometheus.Gauge
}

func NewMetricsHttp(reg prometheus.Registerer) *MetricsHttp {
	m := &MetricsHttp{
		count: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "http",
			Name:      "http_request_count",
			Help:      "Number of http requests.",
		}),
	}
	m.count.Set(float64(0))
	reg.MustRegister(m.count)
	return m
}

func (metric *MetricsHttp) Inc() {
	metric.count.Add(float64(1))
}
