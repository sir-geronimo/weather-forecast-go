package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"weather-api/v1/weather"
)

var (
	router = gin.Default()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := router.Group("/api")

	v1 := api.Group("/v1")
	{
		weather.New(v1)
	}

	defer log.Fatal(router.Run(":8080"))
}
