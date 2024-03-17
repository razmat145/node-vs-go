package controllers

import (
	"fmt"
	"net/http"
)

func FactorialHandler(w http.ResponseWriter, r *http.Request) {
	factorial := computeFactorial(50)

	fmt.Fprint(w, factorial)
}

func computeFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * computeFactorial(n-1)
}
