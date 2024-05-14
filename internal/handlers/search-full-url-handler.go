package handlers

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

func SearchFullURLHandler(writer http.ResponseWriter, request *http.Request, stor storage.URLStogare) {
	shortKey := chi.URLParam(request, "key")
	isFound, fullURL := app.TryGetFullURL([]byte(shortKey), stor)
	if isFound {
		writer.Header().Add("Location", fullURL)
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(writer, "", http.StatusNotFound)
	}
}
