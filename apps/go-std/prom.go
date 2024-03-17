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
)

func init() {
	prometheus.MustRegister(httpDuration)
}

func applyPrometheus(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/metrics" {
			timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path, r.Method))
			defer timer.ObserveDuration()
		}

		next.ServeHTTP(w, r)
	})
}
