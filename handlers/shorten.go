package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
	"url-shortener/utils"
)

func CreateUrl(url string) string {	 // create a new url(call the hash vaala fn) and store it in the database(local map)
	shortUrl := utils.HashUrl(url)
	ctx := context.Background()
	err := utils.RedisClient.Set(ctx, shortUrl, url, 24*time.Hour).Err()

	if err != nil {
		fmt.Println("Error storing URL in Redis:", err)
		return ""
	}


	
	// id := shortUrl
	// models.UrlDb[id] = models.Url{
	// 	Id: id,
	// 	OriginalUrl: url,
	// 	ShortenedUrl: shortUrl,
	// 	CreationDate: time.Now(),
	// }
	return shortUrl
}


func ShortUrlHandler(w http.ResponseWriter, r *http.Request){	// w is used to write responses back by appending to the header
	var data struct {	// struct to hold the data from the request body
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Request!", http.StatusBadRequest)
		return
	}
	shortUrl := CreateUrl(data.URL)
	if shortUrl == ""{
		http.Error(w, "Error storing in redis", http.StatusInternalServerError)
	}

	// response := map[string]string{"shortened_url": shortUrl}
	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Error generating response", http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(jsonResponse)
	
	response := struct {
		ShortUrl string `json:"shortened_url"`
	}{ShortUrl: shortUrl}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error generating Json Response", http.StatusBadRequest)
	}
	ctx := context.Background()
	dbs := utils.RedisClient.Keys(ctx, "*")
	log.Printf("%s Url successfully shortened", data.URL)
	fmt.Println("Shortner In-Memory Content: \n ", dbs)
}

