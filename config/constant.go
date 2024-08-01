package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MONGO_URL string
var TOKEN_HOUR_LIFESPAN string
var API_SECRET string
var REDIS_URL string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	MONGO_URL = os.Getenv("MONGODB_URL")
	REDIS_URL = os.Getenv("REDIS_URL")
	TOKEN_HOUR_LIFESPAN = os.Getenv("TOKEN_HOUR_LIFESPAN")
	API_SECRET = os.Getenv("API_SECRET")
}
