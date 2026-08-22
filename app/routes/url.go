package routes

import (
	"github.com/go-chi/chi/v5"
	handler "main.go/app/handlers"
)

func ApiRouters(r chi.Router) {
	// API Endpoints
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", handler.ServerHealth)
		r.Post("/url", handler.GenerateShortCode)
	})

	// Redirect Endpoint (Supports both GET and HEAD for curl -I / browsers)
	r.Get("/{code}", handler.RedirectURL)
	r.Head("/{code}", handler.RedirectURL)
}