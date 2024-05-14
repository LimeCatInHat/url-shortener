package handlers

import (
	"io"
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
	"github.com/LimeCatInHat/url-shortener/internal/storage"
)

func URLShorterHandler(writer http.ResponseWriter, request *http.Request, storage storage.URLStogare) {
	if request.ContentLength == 0 {
		http.Error(writer, "Invalid Content Length", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	urlResult, err := app.ShortenURL(body, storage)
	if err != nil {
		http.Error(writer, "Attempt to generate short link failed", http.StatusInternalServerError)
	}

	writer.Header().Add("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(urlResult))
}
