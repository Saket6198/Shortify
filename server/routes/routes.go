package routes

import (
	"net/http"
	"github.com/saket6198/url-shortener/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/shorten", handlers.ShortUrlHandler)
	http.HandleFunc("/redirect/", handlers.RedirectUrlHandler)
	http.HandleFunc("/clear", handlers.ClearHandler)
}
