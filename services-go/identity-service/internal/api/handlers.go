package api

import (
	"net/http"
	"github.com/pas-platform/identity-service/internal/domain"
	"github.com/pas-platform/identity-service/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
)

type APIHandler struct {
	storage storage.Storage
}

func NewAPIHandler(s storage.Storage) *APIHandler {
	return &APIHandler{storage: s}
}

func (h *APIHandler) Register(c *gin.Context) {
	var req domain.RegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, tenantID, err := h.storage.CreateTenantAndUser(c.Request.Context(), req)
	if err != nil {
		if _, ok := err.(*storage.ErrDuplicateEmail); ok {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		log.Printf("ERROR: Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process registration"})
		return
	}

	log.Printf("Successfully registered user %s for tenant %s", userID, tenantID)
	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful", "userID": userID, "tenantID": tenantID})
}

func (h *APIHandler) GetToken(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{ "message": "Token endpoint not implemented" })
}