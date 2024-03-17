package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/razmat145/learning/go-echo/controllers"
)

func main() {
	env := newEnv()

	e := echo.New()

	applyPrometheus(e)

	// Routes
	e.GET("/hello", controllers.HelloThere)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", env.Port)))
}
