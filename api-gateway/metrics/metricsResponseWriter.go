package metrics

import "net/http"

type MetricsResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewMetricsResponseWriter(w http.ResponseWriter) *MetricsResponseWriter {
	return &MetricsResponseWriter{w, http.StatusOK}
}

func (metricsWriter *MetricsResponseWriter) WriteHeader(code int) {
	metricsWriter.StatusCode = code
	metricsWriter.ResponseWriter.WriteHeader(code)
}
