package handlers

import "net/http"

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		UrlShorterHandler(writer, request)
	} else {
		SearchFullUrlHandler(writer, request)
	}
}
