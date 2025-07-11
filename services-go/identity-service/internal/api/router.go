package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *APIHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP", "service": "identity-service"})
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/register", handler.Register)
		apiV1.POST("/token", handler.GetToken)
	}

	return r
}