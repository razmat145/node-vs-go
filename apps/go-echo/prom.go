package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func applyPrometheus(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() != "/metrics" {
				timer := prometheus.NewTimer(httpDuration.WithLabelValues(c.Path(), c.Request().Method))
				defer timer.ObserveDuration()
			}
			return next(c)
		}
	})

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
