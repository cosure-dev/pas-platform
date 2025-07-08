package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a new Gin router
	router := gin.Default()

	// Define a simple health check endpoint
	// This tells us the service is up and reachable
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"service": "policy-service",
		})
	})

	// Start the server on port 8080
	// This is the command that keeps the container running
	router.Run(":8080")
}