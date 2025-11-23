// cmd/api/main.go
package main

import (
	"log"

	"shopping-backend/internal/database"
	"shopping-backend/internal/http/router"
)

func main() {
	database.Connect()

	r := router.SetupRouter()

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
