// internal/http/router/router.go
package router

import (
	"github.com/gin-gonic/gin"

	"shopping-backend/internal/http/handlers"
	"shopping-backend/internal/http/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Enable CORS for all routes
	r.Use(middleware.CORSMiddleware())

	// Public routes
	r.POST("/users", handlers.CreateUserHandler)
	r.GET("/users", handlers.ListUsersHandler)
	r.POST("/users/login", handlers.LoginHandler)

	r.POST("/items", handlers.CreateItemHandler)
	r.GET("/items", handlers.ListItemsHandler)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/carts", handlers.CreateOrAddCartHandler)
		auth.GET("/carts", handlers.ListCartsHandler)

		auth.POST("/orders", handlers.CreateOrderHandler)
		auth.GET("/orders", handlers.ListOrdersHandler)
	}

	return r
}
