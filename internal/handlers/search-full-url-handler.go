package handlers

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
	"github.com/go-chi/chi/v5"
)

func SearchFullURLHandler(writer http.ResponseWriter, request *http.Request, stor app.URLStorage) {
	shortKey := chi.URLParam(request, "key")
	fullURL, err := app.GetFullURL([]byte(shortKey), stor)
	if err == nil {
		writer.Header().Add("Location", fullURL)
		writer.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(writer, "", http.StatusNotFound)
	}
}
