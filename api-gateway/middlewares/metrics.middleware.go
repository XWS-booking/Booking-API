package middlewares

import (
	"gateway/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
	"strings"
)

type MetricsMiddleware struct {
	httpCounter *prometheus.CounterVec
	userCounter *prometheus.CounterVec
}

func NewMetricsMiddleware(httpCounter, userCounter *prometheus.CounterVec) MetricsMiddleware {
	return MetricsMiddleware{
		httpCounter: httpCounter,
		userCounter: userCounter,
	}
}

func (middleware *MetricsMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricsWriter := &metrics.MetricsResponseWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(metricsWriter, r)
		status := strconv.Itoa(metricsWriter.StatusCode)
		middleware.httpCounter.WithLabelValues(status).Inc()
		user := strings.Split(r.RemoteAddr, ":")[0]
		middleware.userCounter.WithLabelValues(user).Inc()
	})
}
