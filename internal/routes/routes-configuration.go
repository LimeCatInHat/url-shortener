package routes

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func ConfigureRouter() http.Handler {
	r := chi.NewMux()
	r.Use(middleware.URLFormat)
	r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		handlers.URLShorterHandler(writer, request)
	})
	r.Get("/{key}", func(writer http.ResponseWriter, request *http.Request) {
		handlers.SearchFullURLHandler(writer, request)
	})
	return r
}
