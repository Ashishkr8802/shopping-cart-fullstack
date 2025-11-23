// cmd/api/main.go
package main

import (
    "log"
    "os"

    "shopping-backend/internal/database"
    "shopping-backend/internal/http/router"
)

func main() {
    database.Connect()

    r := router.SetupRouter()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("Server running on port", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}
