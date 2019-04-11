package main

import (
	"net/http"

	"./handlers/basic"
	"./handlers/stackstorm"
)

func newRouter() *http.ServeMux {
	router := http.NewServeMux()

	// routing
	router.Handle("/", basic.Index())
	router.Handle("/healthz", basic.Healthz(&healthy))
	router.Handle("/stackstorm", stackstorm.TriggerSt2Rule(Log))
	// Add more routes here
	// 1. Create a new handler, for example: handlers/new/new.go
	// 2. import "./handlers/new"
	// 3. router.HandleFunc("/new", new.NewHandle)

	return router
}
