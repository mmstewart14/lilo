package domain

import (
	"time"
)

// ClothingItem represents a clothing item in a user's wardrobe
type ClothingItem struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userId"`
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Subcategory string   `json:"subcategory"`
	Color      string    `json:"color"`
	Season     []string  `json:"season"`
	Brand      string    `json:"brand,omitempty"`
	Size       string    `json:"size,omitempty"`
	ImageURLs  []string  `json:"imageUrls"`
	IsOwned    bool      `json:"isOwned"` // true for owned, false for wishlist
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// ClothingCategory represents a category of clothing items
type ClothingCategory struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Subcategories []string `json:"subcategories"`
}

// WardrobeRepository defines the interface for wardrobe data operations
type WardrobeRepository interface {
	CreateItem(item *ClothingItem) error
	GetItemByID(id string) (*ClothingItem, error)
	GetItemsByUserID(userID string, filters map[string]interface{}) ([]*ClothingItem, error)
	UpdateItem(item *ClothingItem) error
	DeleteItem(id string) error
	GetCategories() ([]*ClothingCategory, error)
}

// WardrobeService defines the interface for wardrobe business logic
type WardrobeService interface {
	AddItem(item *ClothingItem) error
	GetItem(id string) (*ClothingItem, error)
	GetUserItems(userID string, filters map[string]interface{}) ([]*ClothingItem, error)
	UpdateItem(item *ClothingItem) error
	DeleteItem(id string) error
	GetCategories() ([]*ClothingCategory, error)
}