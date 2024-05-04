package handlers

import (
	"net/http"
	"strings"

	"github.com/LimeCatInHat/url-shortener/internal/app"
)

func SearchFullURLHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Invalid Http Method", http.StatusBadRequest)
	}
	isFound, shortKey := tryGetShortKeySegment(request)
	if isFound {
		isFound, fullURL := app.TryGetFullURL([]byte(shortKey))
		if isFound {
			writer.Header().Add("Location", fullURL)
			writer.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
	}

	http.Error(writer, "", http.StatusBadRequest)
}

func tryGetShortKeySegment(request *http.Request) (bool, string) {
	parts := strings.Split(strings.Trim(request.URL.Path, "/"), "/")
	return len(parts) == 1, parts[0]
}
