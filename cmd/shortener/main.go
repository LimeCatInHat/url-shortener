package main

import (
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/config"
	"github.com/LimeCatInHat/url-shortener/internal/routes"
)

func main() {
	config.SetConfiguration()
	r := routes.ConfigureRouter()

	err := http.ListenAndServe(config.AppSettings.SrvAddr, r)
	if err != nil {
		panic(err)
	}
}
