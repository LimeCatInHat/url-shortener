package handlers

import (
	"io"
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/app"
)

type requestValidationResult struct {
	isValid            bool
	message            string
	responseStatusCode int
}

func UrlShorterHandler(writer http.ResponseWriter, request *http.Request) {
	validationResult := validateShorterHandlerRequest(request)
	if !validationResult.isValid {
		http.Error(writer, validationResult.message, validationResult.responseStatusCode)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	urlResult := app.ShortenUrl(body)

	writer.Header().Add("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(urlResult))
}

func validateShorterHandlerRequest(request *http.Request) requestValidationResult {
	if request.Method != http.MethodPost {
		return requestValidationResult{message: "Invalid Http Method", responseStatusCode: http.StatusMethodNotAllowed}
	}
	if request.Header.Get("Content-Type") != "text/plain" {
		return requestValidationResult{message: "Invalid Content-Type", responseStatusCode: http.StatusBadRequest}
	}
	if request.ContentLength == 0 {
		return requestValidationResult{message: "Invalid Content Length", responseStatusCode: http.StatusBadRequest}
	}

	return requestValidationResult{isValid: true}
}
