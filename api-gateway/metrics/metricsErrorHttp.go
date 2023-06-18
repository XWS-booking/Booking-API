package metrics

import "github.com/prometheus/client_golang/prometheus"

type MetricsErrorHttp struct {
	count prometheus.Gauge
}

func NewMetricsHttpError(reg prometheus.Registerer) *MetricsErrorHttp {
	m := &MetricsErrorHttp{
		count: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "http",
			Name:      "http_error_request_count",
			Help:      "Number of http requests.",
		}),
	}
	m.count.Set(float64(0))
	reg.MustRegister(m.count)
	return m
}

func (metric *MetricsErrorHttp) Inc() {
	metric.count.Add(float64(1))
}
