package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"url-shortener/utils"
)

func Geturl(id string) (string, error){
	// url, exists := models.UrlDb[id]
	ctx := context.Background()
	url, err := utils.RedisClient.Get(ctx, id).Result()
	if err != nil {
		return "", fmt.Errorf("error retrieving URL from Redis: %v", err)
	}
	return url, nil

	// if !exists {
	// 	return models.Url{}, fmt.Errorf("url not found")
	// }
	// return url, nil

}

func RedirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := Geturl(id)
	if err != nil {
		http.Error(w, "Url not found", http.StatusNotFound)
		return
	}
	log.Printf("Redirecting to %s", url)
	http.Redirect(w, r, url, http.StatusFound)

}