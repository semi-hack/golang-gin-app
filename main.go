package main

import (
	"gin/db"
	"gin/routes"
	"github.com/joho/godotenv"
	"log"
	// "github.com/gin-gonic/gin"
)

func init() {
    // Log error if .env file does not exist
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }
}

func main() {
	db.GetConnection()
	routes.Initialize()
}
