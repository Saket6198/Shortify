package handlers

import (
	"context"
	"log"
	"fmt"
	"net/http"
	"github.com/saket6198/url-shortener/utils"
)

func ClearHandler(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	err := utils.RedisClient.FlushAll(ctx).Err()
	if err != nil {
		log.Printf("Error clearing Redis: %v", err)
		http.Error(w, "Error clearing Redis", http.StatusInternalServerError)
		return
	}
	log.Println("Cleared all data from Redis")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Cleared all data from Redis")
}
