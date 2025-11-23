// internal/models/cart.go
package models

import "time"

type Cart struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UserID    uint       `json:"user_id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"` // open/ordered
	CreatedAt time.Time  `json:"created_at"`
	Items     []CartItem `json:"items"`
}

type CartItem struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CartID    uint      `json:"cart_id"`
	ItemID    uint      `json:"item_id"`
	CreatedAt time.Time `json:"created_at"`
	Item      Item      `json:"item"`
}
