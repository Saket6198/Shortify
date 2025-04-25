package utils

import (
	"context"
	"fmt"
	"log"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client	// create a global obj of the class redis

func InitRedis(){
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-16367.crce182.ap-south-1-1.ec2.redns.redis-cloud.com:16367",
		Username: "default",
		Password: "Rt3s9TPh2Yu90anDqsRa5TrApmInv3SX",
		DB:       0,
	})
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	} else {
		fmt.Println("Connected to Redis")
	}
}

