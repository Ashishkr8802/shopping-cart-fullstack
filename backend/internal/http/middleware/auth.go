// internal/http/middleware/auth.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shopping-backend/internal/database"
	"shopping-backend/internal/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		var user models.User
		if err := database.DB.Where("token = ?", token).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func CurrentUser(c *gin.Context) models.User {
	u, _ := c.Get("user")
	return u.(models.User)
}
