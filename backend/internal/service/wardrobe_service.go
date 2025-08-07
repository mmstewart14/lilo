package service

import (
	"errors"
	"fmt"

	"github.com/lilo/backend/internal/domain"
)

// WardrobeServiceImpl implements WardrobeService
type WardrobeServiceImpl struct {
	wardrobeRepo domain.WardrobeRepository
}

// NewWardrobeService creates a new wardrobe service
func NewWardrobeService(wardrobeRepo domain.WardrobeRepository) domain.WardrobeService {
	return &WardrobeServiceImpl{
		wardrobeRepo: wardrobeRepo,
	}
}

// AddItem adds a new clothing item to the wardrobe
func (s *WardrobeServiceImpl) AddItem(item *domain.ClothingItem) error {
	// Validate required fields
	if item.UserID == "" {
		return errors.New("user ID is required")
	}
	if item.Name == "" {
		return errors.New("item name is required")
	}
	if item.Category == "" {
		return errors.New("category is required")
	}
	if item.Color == "" {
		return errors.New("color is required")
	}

	// Set default values if not provided
	if len(item.Season) == 0 {
		item.Season = []string{"Spring", "Summer", "Fall", "Winter"}
	}
	if len(item.ImageURLs) == 0 {
		item.ImageURLs = []string{}
	}

	return s.wardrobeRepo.CreateItem(item)
}

// GetItem retrieves a clothing item by ID
func (s *WardrobeServiceImpl) GetItem(id string) (*domain.ClothingItem, error) {
	if id == "" {
		return nil, errors.New("item ID is required")
	}
	return s.wardrobeRepo.GetItemByID(id)
}

// GetUserItems retrieves all clothing items for a user with optional filters
func (s *WardrobeServiceImpl) GetUserItems(userID string, filters map[string]interface{}) ([]*domain.ClothingItem, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	return s.wardrobeRepo.GetItemsByUserID(userID, filters)
}

// UpdateItem updates an existing clothing item
func (s *WardrobeServiceImpl) UpdateItem(item *domain.ClothingItem) error {
	if item.ID == "" {
		return errors.New("item ID is required")
	}
	if item.UserID == "" {
		return errors.New("user ID is required")
	}
	if item.Name == "" {
		return errors.New("item name is required")
	}
	if item.Category == "" {
		return errors.New("category is required")
	}
	if item.Color == "" {
		return errors.New("color is required")
	}

	// Verify item exists
	existingItem, err := s.wardrobeRepo.GetItemByID(item.ID)
	if err != nil {
		return fmt.Errorf("item not found: %w", err)
	}

	// Verify user owns the item
	if existingItem.UserID != item.UserID {
		return errors.New("unauthorized: item belongs to different user")
	}

	return s.wardrobeRepo.UpdateItem(item)
}

// DeleteItem deletes a clothing item by ID
func (s *WardrobeServiceImpl) DeleteItem(id string) error {
	if id == "" {
		return errors.New("item ID is required")
	}

	// Verify item exists before deletion
	_, err := s.wardrobeRepo.GetItemByID(id)
	if err != nil {
		return fmt.Errorf("item not found: %w", err)
	}

	return s.wardrobeRepo.DeleteItem(id)
}

// GetCategories retrieves all available clothing categories
func (s *WardrobeServiceImpl) GetCategories() ([]*domain.ClothingCategory, error) {
	return s.wardrobeRepo.GetCategories()
}
