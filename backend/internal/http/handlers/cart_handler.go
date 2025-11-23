// internal/http/handlers/cart_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-backend/internal/database"
	"shopping-backend/internal/http/middleware"
	"shopping-backend/internal/models"
)

type CartRequest struct {
	Items []uint `json:"items" binding:"required"`
}

// POST /carts
func CreateOrAddCartHandler(c *gin.Context) {
	user := middleware.CurrentUser(c)

	var req CartRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "items array is required"})
		return
	}

	var cart models.Cart
	if err := database.DB.
		Where("user_id = ? AND status = ?", user.ID, "open").
		First(&cart).Error; err != nil {

		cart = models.Cart{
			UserID: user.ID,
			Name:   "Cart for " + user.Username,
			Status: "open",
		}
		if err := database.DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create cart"})
			return
		}
		user.CartID = &cart.ID
		database.DB.Save(&user)
	}

	for _, itemID := range req.Items {
		ci := models.CartItem{
			CartID: cart.ID,
			ItemID: itemID,
		}
		database.DB.Create(&ci)
	}

	var items []models.CartItem
	database.DB.Preload("Item").Where("cart_id = ?", cart.ID).Find(&items)
	cart.Items = items

	c.JSON(http.StatusOK, cart)
}

// GET /carts
func ListCartsHandler(c *gin.Context) {
	var carts []models.Cart
	database.DB.Preload("Items").Find(&carts)
	c.JSON(http.StatusOK, carts)
}
