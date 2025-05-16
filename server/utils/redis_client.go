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
		Addr:     "redis-11272.c301.ap-south-1-1.ec2.redns.redis-cloud.com:11272",
		Username: "default",
		Password: "WVmmav4q4muocgPKHQMd9HibmNGrIq9e",
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

