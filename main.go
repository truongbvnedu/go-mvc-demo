package main

import (
	"go-mvc-demo/config"
	routes "go-mvc-demo/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ConnectDatabase()
	r := routes.SetupRouter()
	port := os.Getenv("PORT")

	r.Run(":" + port)
}
