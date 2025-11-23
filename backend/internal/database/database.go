// internal/database/database.go
package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"shopping-backend/internal/models"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open("sqlite3", "shopping.db")
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	DB.LogMode(true)

	// Auto-migrate models
	DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
	)
}
