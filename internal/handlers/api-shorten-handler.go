package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
	"github.com/LimeCatInHat/url-shortener/internal/models"
)

func APIShorterHandler(writer http.ResponseWriter, request *http.Request, stor app.URLStorage) {
	if request.ContentLength == 0 {
		http.Error(writer, "invalid content length", http.StatusBadRequest)
		return
	}

	var req models.ShortenURLRequest
	dec := json.NewDecoder(request.Body)
	if err := dec.Decode(&req); err != nil {
		log.Printf(`cannot decode request JSON body: %v`, err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	urlResult, err := app.ShortenURL([]byte(req.URL), stor)
	if err != nil {
		log.Printf(`attempt to generate short link for '%q' failed: %v`, req.URL, err)
		http.Error(writer, "attempt to generate short link failed", http.StatusInternalServerError)
		return
	}

	resp := models.ShortenURLResponse{
		ShortLink: urlResult,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(writer)
	if err := enc.Encode(resp); err != nil {
		http.Error(writer, "attempt to encode response failed", http.StatusInternalServerError)
		return
	}
}
