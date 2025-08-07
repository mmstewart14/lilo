package repository

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lilo/backend/internal/domain"
)

// InMemoryWardrobeRepository implements WardrobeRepository using in-memory storage
type InMemoryWardrobeRepository struct {
	items      map[string]*domain.ClothingItem
	categories []*domain.ClothingCategory
	mu         sync.RWMutex
}

// NewWardrobeRepository creates a new wardrobe repository
func NewWardrobeRepository() domain.WardrobeRepository {
	repo := &InMemoryWardrobeRepository{
		items: make(map[string]*domain.ClothingItem),
	}

	// Initialize default categories
	repo.initializeCategories()

	return repo
}

// initializeCategories sets up default clothing categories
func (r *InMemoryWardrobeRepository) initializeCategories() {
	r.categories = []*domain.ClothingCategory{
		{
			ID:            "tops",
			Name:          "Tops",
			Subcategories: []string{"T-Shirts", "Shirts", "Blouses", "Sweaters", "Hoodies", "Tank Tops"},
		},
		{
			ID:            "bottoms",
			Name:          "Bottoms",
			Subcategories: []string{"Jeans", "Pants", "Shorts", "Skirts", "Leggings"},
		},
		{
			ID:            "dresses",
			Name:          "Dresses",
			Subcategories: []string{"Casual Dresses", "Formal Dresses", "Maxi Dresses", "Mini Dresses"},
		},
		{
			ID:            "outerwear",
			Name:          "Outerwear",
			Subcategories: []string{"Jackets", "Coats", "Blazers", "Cardigans", "Vests"},
		},
		{
			ID:            "shoes",
			Name:          "Shoes",
			Subcategories: []string{"Sneakers", "Boots", "Heels", "Flats", "Sandals", "Athletic Shoes"},
		},
		{
			ID:            "accessories",
			Name:          "Accessories",
			Subcategories: []string{"Bags", "Jewelry", "Hats", "Scarves", "Belts", "Watches"},
		},
	}
}

// CreateItem creates a new clothing item
func (r *InMemoryWardrobeRepository) CreateItem(item *domain.ClothingItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if item.ID == "" {
		item.ID = uuid.New().String()
	}
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	r.items[item.ID] = item
	return nil
}

// GetItemByID retrieves a clothing item by ID
func (r *InMemoryWardrobeRepository) GetItemByID(id string) (*domain.ClothingItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.items[id]
	if !exists {
		return nil, domain.ErrUserNotFound // Reusing error for simplicity
	}
	return item, nil
}

// GetItemsByUserID retrieves all clothing items for a user with optional filters
func (r *InMemoryWardrobeRepository) GetItemsByUserID(userID string, filters map[string]interface{}) ([]*domain.ClothingItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var items []*domain.ClothingItem
	for _, item := range r.items {
		if item.UserID == userID {
			// Apply filters if provided
			if r.matchesFilters(item, filters) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

// matchesFilters checks if an item matches the provided filters
func (r *InMemoryWardrobeRepository) matchesFilters(item *domain.ClothingItem, filters map[string]interface{}) bool {
	if filters == nil {
		return true
	}

	if category, ok := filters["category"]; ok {
		if item.Category != category.(string) {
			return false
		}
	}

	if color, ok := filters["color"]; ok {
		if item.Color != color.(string) {
			return false
		}
	}

	if isOwned, ok := filters["isOwned"]; ok {
		if item.IsOwned != isOwned.(bool) {
			return false
		}
	}

	if season, ok := filters["season"]; ok {
		seasonStr := season.(string)
		found := false
		for _, s := range item.Season {
			if s == seasonStr {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// UpdateItem updates an existing clothing item
func (r *InMemoryWardrobeRepository) UpdateItem(item *domain.ClothingItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.items[item.ID]; !exists {
		return domain.ErrUserNotFound // Reusing error for simplicity
	}

	item.UpdatedAt = time.Now()
	r.items[item.ID] = item
	return nil
}

// DeleteItem deletes a clothing item by ID
func (r *InMemoryWardrobeRepository) DeleteItem(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.items[id]; !exists {
		return domain.ErrUserNotFound // Reusing error for simplicity
	}

	delete(r.items, id)
	return nil
}

// GetCategories retrieves all clothing categories
func (r *InMemoryWardrobeRepository) GetCategories() ([]*domain.ClothingCategory, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.categories, nil
}
