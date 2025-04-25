package main

import (
	"fmt"
	"log"
	"github.com/saket6198/url-shortener/utils"
	"net/http"
	"github.com/saket6198/url-shortener/routes"
	"github.com/rs/cors"
)


func main(){
	fmt.Println("Starting Url Shortener...")
	fmt.Println("Server starting on port 5000")
	utils.InitRedis()
	routes.RegisterRoutes()

	handler:= cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(http.DefaultServeMux)


	err := http.ListenAndServe(":5000", handler); if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}