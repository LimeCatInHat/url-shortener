package handlers

import "net/http"

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		URLShorterHandler(writer, request)
	} else {
		SearchFullURLHandler(writer, request)
	}
}
