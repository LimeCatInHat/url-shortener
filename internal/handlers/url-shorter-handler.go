package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
)

func URLShorterHandler(writer http.ResponseWriter, request *http.Request, stor app.URLStorage) {
	if request.ContentLength == 0 {
		http.Error(writer, "Invalid Content Length", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	urlResult, err := app.ShortenURL(body, stor)
	if err != nil {
		log.Printf(`attempt to generate short link for '%q' failed: %v`, body, err)
		http.Error(writer, "attempt to generate short link failed", http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte(urlResult))
	if err != nil {
		http.Error(writer, "attempt to send short link failed", http.StatusInternalServerError)
		return
	}
}
