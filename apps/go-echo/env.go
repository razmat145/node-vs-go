package main

import "os"

type env struct {
	Port string
}

func (e *env) load() {
	port, isSet := os.LookupEnv("API_PORT")

	if !isSet {
		port = "8989"
	}

	e.Port = port
}

func newEnv() *env {
	e := env{}
	e.load()
	return &e
}
