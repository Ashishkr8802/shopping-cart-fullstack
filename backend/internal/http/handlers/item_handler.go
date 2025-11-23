// internal/http/handlers/item_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-backend/internal/database"
	"shopping-backend/internal/models"
)

type CreateItemRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status"`
}

// POST /items
func CreateItemHandler(c *gin.Context) {
	var req CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{
		Name:   req.Name,
		Status: req.Status,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GET /items
func ListItemsHandler(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}
