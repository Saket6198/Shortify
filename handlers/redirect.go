package handlers

import (
	"net/http"
	"log"	
	"fmt"
	"url-shortener/models"
	
)

func Geturl(id string) (models.Url, error){
	url, exists := models.UrlDb[id]
	if !exists {
		return models.Url{}, fmt.Errorf("Url not found")
	}
	return url, nil

}

func RedirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := Geturl(id)
	if err != nil {
		http.Error(w, "Url not found", http.StatusNotFound)
		return
	}
	log.Printf("Redirecting to %s", url.OriginalUrl)
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)

}