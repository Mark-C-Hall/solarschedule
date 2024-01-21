package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mark-c-hall/solarschedule/internal/api"
)

func main() {
	router := gin.Default()

	// Setup Route
	api.SetupRoutes(router)

	// Start Server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
