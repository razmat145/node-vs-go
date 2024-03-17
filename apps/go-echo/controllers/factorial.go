package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Factorial(c echo.Context) error {
	factorial := computeFactorial(100)

	return c.String(http.StatusOK, strconv.Itoa(factorial))
}

func computeFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * computeFactorial(n-1)
}
