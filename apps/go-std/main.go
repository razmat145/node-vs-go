package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/razmat145/learning/go-std/controllers"
)

func main() {
	env := newEnv()

	http.Handle("/hello", applyPrometheus(controllers.HelloHandler))
	http.Handle("/factorial", applyPrometheus(controllers.FactorialHandler))
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Starting server at port %s\n", env.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", env.Port), nil)
}
