package routes

import (
	"net/http"
	"url-shortener/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/shorten", handlers.ShortUrlHandler)
	http.HandleFunc("/redirect/", handlers.RedirectUrlHandler)
	http.HandleFunc("/clear", handlers.ClearHandler)
}
