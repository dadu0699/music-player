package main

import (
	"api/pkg/app"
	"api/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func run() error {
	database.Init()

	router := gin.Default()
	router.Use(cors.Default())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	app.Routes(router)

	if err := router.Run(); err != nil {
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
