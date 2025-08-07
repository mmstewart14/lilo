package service

import (
	"errors"
	"fmt"

	"github.com/lilo/backend/internal/domain"
)

// OutfitServiceImpl implements OutfitService
type OutfitServiceImpl struct {
	outfitRepo domain.OutfitRepository
}

// NewOutfitService creates a new outfit service
func NewOutfitService(outfitRepo domain.OutfitRepository) domain.OutfitService {
	return &OutfitServiceImpl{
		outfitRepo: outfitRepo,
	}
}

// CreateOutfit creates a new outfit
func (s *OutfitServiceImpl) CreateOutfit(outfit *domain.Outfit) error {
	// Validate required fields
	if outfit.UserID == "" {
		return errors.New("user ID is required")
	}
	if outfit.Name == "" {
		return errors.New("outfit name is required")
	}
	if len(outfit.Items) == 0 {
		return errors.New("outfit must contain at least one item")
	}

	// Set default values if not provided
	if len(outfit.Occasion) == 0 {
		outfit.Occasion = []string{"casual"}
	}
	if len(outfit.Season) == 0 {
		outfit.Season = []string{"Spring", "Summer", "Fall", "Winter"}
	}

	return s.outfitRepo.CreateOutfit(outfit)
}

// GetOutfit retrieves an outfit by ID
func (s *OutfitServiceImpl) GetOutfit(id string) (*domain.Outfit, error) {
	if id == "" {
		return nil, errors.New("outfit ID is required")
	}
	return s.outfitRepo.GetOutfitByID(id)
}

// GetUserOutfits retrieves all outfits for a user with optional filters
func (s *OutfitServiceImpl) GetUserOutfits(userID string, filters map[string]interface{}) ([]*domain.Outfit, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	return s.outfitRepo.GetOutfitsByUserID(userID, filters)
}

// UpdateOutfit updates an existing outfit
func (s *OutfitServiceImpl) UpdateOutfit(outfit *domain.Outfit) error {
	if outfit.ID == "" {
		return errors.New("outfit ID is required")
	}
	if outfit.UserID == "" {
		return errors.New("user ID is required")
	}
	if outfit.Name == "" {
		return errors.New("outfit name is required")
	}
	if len(outfit.Items) == 0 {
		return errors.New("outfit must contain at least one item")
	}

	// Verify outfit exists
	existingOutfit, err := s.outfitRepo.GetOutfitByID(outfit.ID)
	if err != nil {
		return fmt.Errorf("outfit not found: %w", err)
	}

	// Verify user owns the outfit
	if existingOutfit.UserID != outfit.UserID {
		return errors.New("unauthorized: outfit belongs to different user")
	}

	return s.outfitRepo.UpdateOutfit(outfit)
}

// DeleteOutfit deletes an outfit by ID
func (s *OutfitServiceImpl) DeleteOutfit(id string) error {
	if id == "" {
		return errors.New("outfit ID is required")
	}

	// Verify outfit exists before deletion
	_, err := s.outfitRepo.GetOutfitByID(id)
	if err != nil {
		return fmt.Errorf("outfit not found: %w", err)
	}

	return s.outfitRepo.DeleteOutfit(id)
}

// FavoriteOutfit marks an outfit as favorite
func (s *OutfitServiceImpl) FavoriteOutfit(id string) error {
	if id == "" {
		return errors.New("outfit ID is required")
	}

	// Verify outfit exists
	_, err := s.outfitRepo.GetOutfitByID(id)
	if err != nil {
		return fmt.Errorf("outfit not found: %w", err)
	}

	return s.outfitRepo.SetFavorite(id, true)
}

// UnfavoriteOutfit removes favorite status from an outfit
func (s *OutfitServiceImpl) UnfavoriteOutfit(id string) error {
	if id == "" {
		return errors.New("outfit ID is required")
	}

	// Verify outfit exists
	_, err := s.outfitRepo.GetOutfitByID(id)
	if err != nil {
		return fmt.Errorf("outfit not found: %w", err)
	}

	return s.outfitRepo.SetFavorite(id, false)
}

// SubmitReflection submits a reflection for an outfit
func (s *OutfitServiceImpl) SubmitReflection(reflection *domain.Reflection) error {
	// Validate required fields
	if reflection.UserID == "" {
		return errors.New("user ID is required")
	}
	if reflection.OutfitID == "" {
		return errors.New("outfit ID is required")
	}
	if reflection.Confidence < 1 || reflection.Confidence > 5 {
		return errors.New("confidence must be between 1 and 5")
	}
	if reflection.Comfort < 1 || reflection.Comfort > 5 {
		return errors.New("comfort must be between 1 and 5")
	}

	// Verify outfit exists
	outfit, err := s.outfitRepo.GetOutfitByID(reflection.OutfitID)
	if err != nil {
		return fmt.Errorf("outfit not found: %w", err)
	}

	// Verify user owns the outfit
	if outfit.UserID != reflection.UserID {
		return errors.New("unauthorized: outfit belongs to different user")
	}

	return s.outfitRepo.CreateReflection(reflection)
}

// GetUserReflections retrieves all reflections for a user
func (s *OutfitServiceImpl) GetUserReflections(userID string) ([]*domain.Reflection, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	return s.outfitRepo.GetReflectionsByUserID(userID)
}
