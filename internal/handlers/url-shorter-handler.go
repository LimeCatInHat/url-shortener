package handlers

import (
	"io"
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
)

func URLShorterHandler(writer http.ResponseWriter, request *http.Request) {
	if request.ContentLength == 0 {
		http.Error(writer, "Invalid Content Length", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	urlResult := app.ShortenURL(body)

	writer.Header().Add("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(urlResult))
}
