package main

import (
	"os"
	"strconv"
)

type env struct {
	Concurrency int
	Url         string
}

func (e *env) load() {
	e.loadConcurrency()
	e.loadUrl()
}

func (e *env) loadConcurrency() {
	concurrencyStr, isSet := os.LookupEnv("CONCURRENCY")

	concurrency, err := strconv.Atoi(concurrencyStr)
	if !isSet || err != nil {
		e.Concurrency = 25_000
		return
	}

	e.Concurrency = concurrency
}

func (e *env) loadUrl() {
	url, isSet := os.LookupEnv("SPAM_URL")

	if !isSet {
		panic("URL is not set")
	}

	e.Url = url
}

func newEnv() *env {
	e := env{}

	e.load()

	return &e
}
