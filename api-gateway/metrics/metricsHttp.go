package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type MetricsHttp struct {
	counter *prometheus.CounterVec
}

func NewMetricsHttp(reg prometheus.Registerer, statuses []int) *MetricsHttp {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "http",
			Name:      "number_of_requests",
			Help:      "Number of HTTP requests by status code.",
		},
		[]string{"status"},
	)

	for _, status := range statuses {
		counter.WithLabelValues(strconv.Itoa(status)).Add(0)
	}

	reg.MustRegister(counter)

	return &MetricsHttp{
		counter: counter,
	}
}

func (metric *MetricsHttp) Inc(status int) {
	metric.counter.WithLabelValues(strconv.Itoa(status)).Inc()
}
