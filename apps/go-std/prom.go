package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path", "method"})

	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	}, []string{"path", "method"})
)

func init() {
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(httpRequestsTotal)
}

func applyPrometheus(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path, r.Method))
		defer timer.ObserveDuration()

		httpRequestsTotal.WithLabelValues(r.URL.Path, r.Method).Inc()

		next.ServeHTTP(w, r)
	})
}
