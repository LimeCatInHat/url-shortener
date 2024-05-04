package routes

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/handlers"
)

func ConfigureRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.RootHandler)
}
