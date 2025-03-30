package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"url-shortener/models"
	"url-shortener/utils"

)

func CreateUrl(url string) string {	 // create a new url(call the hash vaala fn) and store it in the database(local map)
	shortUrl := utils.HashUrl(url)
	id := shortUrl
	models.UrlDb[id] = models.Url{
		Id: id,
		OriginalUrl: url,
		ShortenedUrl: shortUrl,
		CreationDate: time.Now(),
	}
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
	fmt.Println("Shortner In-Memory Content: \n ", models.UrlDb)
}

