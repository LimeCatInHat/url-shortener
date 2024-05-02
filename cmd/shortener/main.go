package main

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/configuration"
	"github.com/LimeCatInHat/url-shortener/internal/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.ConfigureRoutes(mux)

	err := http.ListenAndServe(configuration.AppUrl, mux)
	if err != nil {
		panic(err)
	}
}
