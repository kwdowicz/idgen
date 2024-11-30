package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func v1Router(a application) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", a.generate)
	return r
}