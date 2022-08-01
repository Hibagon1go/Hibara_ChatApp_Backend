package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()
var redisClient *redis.Client

func init() {
	err := godotenv.Load("env/.env.dev")
	if err != nil {
		fmt.Printf("Error loading .env.dev file: %v", err)
	}
	redisClient = redis.NewClient(&redis.Options{
		Username: "red-cauf4sl0mal15ku91m20",
		Addr:     "oregon-redis.render.com:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			//Certificates: []tls.Certificate{cert}
		},
	})
}
