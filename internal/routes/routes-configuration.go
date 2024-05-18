package routes

import (
	"fmt"
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
	"github.com/LimeCatInHat/url-shortener/internal/handlers"
	"github.com/LimeCatInHat/url-shortener/internal/logger"
	"github.com/LimeCatInHat/url-shortener/internal/middlewares"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

func ConfigureRouter() (http.Handler, error) {
	r := chi.NewMux()
	requestLogger, err := logger.CreateRequestLogger()
	if err != nil {
		return nil, fmt.Errorf("request logger wasn't created: %w", err)
	}

	var stor = storage.GetStorage()
	r.Post("/", middlewares.WithLogging(shorterHandler(stor), requestLogger))
	r.Get("/{key}", middlewares.WithLogging(urlSearcher(stor), requestLogger))

	return r, nil
}

func shorterHandler(stor app.URLStorage) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		handlers.URLShorterHandler(w, r, stor)
	}
	return http.HandlerFunc(fn)
}

func urlSearcher(stor app.URLStorage) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		handlers.SearchFullURLHandler(w, r, stor)
	}
	return http.HandlerFunc(fn)
}
