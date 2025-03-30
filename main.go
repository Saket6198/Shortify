package main

import (
	"crypto/sha256"
	"encoding/hex"
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

func UnHash(w http.ResponseWriter, r *http.Request){	// w is used to write responses back by appending to the header
	id := 
}

func main(){
	fmt.Println("Starting Url Shortener...")
	fmt.Println("Server starting on port 5000")
	http.HandleFunc("/", example)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
	// fmt.Println("Hashing done", hashUrl("www.google.com"))

}