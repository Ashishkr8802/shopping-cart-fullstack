// internal/http/handlers/order_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-backend/internal/database"
	"shopping-backend/internal/http/middleware"
	"shopping-backend/internal/models"
)

type CreateOrderRequest struct {
	CartID uint `json:"cart_id" binding:"required"`
}

// POST /orders
func CreateOrderHandler(c *gin.Context) {
	user := middleware.CurrentUser(c)

	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("id = ? AND user_id = ?", req.CartID, user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found for user"})
		return
	}

	order := models.Order{
		CartID: cart.ID,
		UserID: user.ID,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create order"})
		return
	}

	cart.Status = "ordered"
	database.DB.Save(&cart)

	c.JSON(http.StatusCreated, order)
}

// GET /orders
func ListOrdersHandler(c *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}
