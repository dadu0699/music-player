package main

import (
	"api/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func run() error {
	database.Init()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s: %s", "Error loading .env file", err)
	}

	if err := run(); err != nil {
		log.Fatalf("This is the startup error: %s", err)
		os.Exit(1)
	}
}
