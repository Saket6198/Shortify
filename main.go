package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Url struct {
	Id string `json:"id`
	OriginalUrl string `json: "original_url"`
	ShortenedUrl string `json: "shortened_url"`
	CreationDate time.Time `json: "creation_date"`

}

var urlDb = make(map[string]Url)

func hashUrl(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashsum := hasher.Sum(nil)
	hashhex := hex.EncodeToString(hashsum)
	return hashhex[:16]
	// fmt.Println("hash: ", hashhex)
	// return "hi"
}

func CreateUrl(url string) string {
	shortUrl := hashUrl(url)
	id := shortUrl[:8]
	urlDb[id] = Url{
		Id: id,
		OriginalUrl: url,
		ShortenedUrl: shortUrl,
		CreationDate: time.Now(),
	}
	return shortUrl
}


func Geturl(id string) (Url, error){
	url, exists := urlDb[id]
	if !exists {
		return Url{}, fmt.Errorf("Url not found")
	}
	return url, nil

}

func example(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "testing")
	fmt.Println("Request received")
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
}


func main(){
	fmt.Println("Starting Url Shortener...")
	fmt.Println("Server starting on port 5000")
	http.HandleFunc("/", example)
	http.HandleFunc("/shorten", ShortUrlHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
	// fmt.Println("Hashing done", hashUrl("www.google.com"))
	
}