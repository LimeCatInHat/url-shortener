package routes

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/handlers"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

func ConfigureRouter() http.Handler {
	r := chi.NewMux()
	stor := storage.GetStorage()
	r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		handlers.URLShorterHandler(writer, request, stor)
	})
	r.Get("/{key}", func(writer http.ResponseWriter, request *http.Request) {
		handlers.SearchFullURLHandler(writer, request, stor)
	})
	return r
}
