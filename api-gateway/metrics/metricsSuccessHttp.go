package metrics

import "github.com/prometheus/client_golang/prometheus"

type MetricsSuccessHttp struct {
	count prometheus.Gauge
}

func NewMetricsHttpSuccess(reg prometheus.Registerer) *MetricsSuccessHttp {
	m := &MetricsSuccessHttp{
		count: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "http",
			Name:      "http_success_request_count",
			Help:      "Number of http requests.",
		}),
	}
	m.count.Set(float64(0))
	reg.MustRegister(m.count)
	return m
}

func (metric *MetricsSuccessHttp) Inc() {
	metric.count.Add(float64(1))
}
