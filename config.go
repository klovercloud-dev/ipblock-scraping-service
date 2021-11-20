package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitDb() *redis.Client {
	loadEnvironments()

	fmt.Println("REDIS_DB: ", os.Getenv("REDIS_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECT_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}

func loadEnvironments() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
