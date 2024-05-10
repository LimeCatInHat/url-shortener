package main

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/configuration"
	"github.com/LimeCatInHat/url-shortener/internal/routes"
)

func main() {
	r := routes.ConfigureRouter()

	err := http.ListenAndServe(configuration.AppURL, r)
	if err != nil {
		panic(err)
	}
}
