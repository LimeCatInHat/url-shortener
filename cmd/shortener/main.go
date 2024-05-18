package main

import (
	"log"
	"net/http"

	"github.com/LimeCatInHat/url-shortener/internal/config"
	"github.com/LimeCatInHat/url-shortener/internal/routes"
)

func main() {
	configuration := config.Init()
	r, err := routes.ConfigureRouter()
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(configuration.SrvAddr, r)

	if err != nil {
		log.Fatal(err)
	}
}
