package main

import (
	"fmt"
	"log"
	"url-shortener/utils"
	"net/http"
	"url-shortener/routes"
)


func main(){
	fmt.Println("Starting Url Shortener...")
	fmt.Println("Server starting on port 5000")
	utils.InitRedis()
	routes.RegisterRoutes()
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}