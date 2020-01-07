package main

import (
	"auth/db"
	"auth/routes"
	// "github.com/gin-gonic/gin"
)

func main() {
	db.GetConnection()
	routes.Initialize()
}
