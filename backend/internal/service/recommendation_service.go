package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/lilo/backend/internal/domain"
)

// RecommendationServiceImpl implements RecommendationService
type RecommendationServiceImpl struct {
	recommendationRepo domain.RecommendationRepository
	wardrobeRepo       domain.WardrobeRepository
	outfitRepo         domain.OutfitRepository
}

// NewRecommendationService creates a new recommendation service
func NewRecommendationService(
	recommendationRepo domain.RecommendationRepository,
	wardrobeRepo domain.WardrobeRepository,
	outfitRepo domain.OutfitRepository,
) domain.RecommendationService {
	return &RecommendationServiceImpl{
		recommendationRepo: recommendationRepo,
		wardrobeRepo:       wardrobeRepo,
		outfitRepo:         outfitRepo,
	}
}

// GetDailyRecommendations generates daily outfit recommendations for a user
func (s *RecommendationServiceImpl) GetDailyRecommendations(userID string) ([]*domain.Outfit, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	// Get user's outfits
	outfits, err := s.outfitRepo.GetOutfitsByUserID(userID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get user outfits: %w", err)
	}

	if len(outfits) == 0 {
		return []*domain.Outfit{}, nil
	}

	// Simple recommendation algorithm: return up to 3 random outfits
	// In a real implementation, this would be much more sophisticated
	rand.Seed(time.Now().UnixNano())

	numRecommendations := 3
	if len(outfits) < numRecommendations {
		numRecommendations = len(outfits)
	}

	// Shuffle and take first N outfits
	shuffled := make([]*domain.Outfit, len(outfits))
	copy(shuffled, outfits)

	for i := range shuffled {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	recommendations := shuffled[:numRecommendations]

	// Mark them as recommended
	for _, outfit := range recommendations {
		outfit.IsRecommended = true
	}

	return recommendations, nil
}

// GetExploreRecommendations generates explore recommendations for a user with filters
func (s *RecommendationServiceImpl) GetExploreRecommendations(userID string, filters map[string]interface{}) ([]*domain.Outfit, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	// Get user's outfits with filters
	outfits, err := s.outfitRepo.GetOutfitsByUserID(userID, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to get user outfits: %w", err)
	}

	// Simple exploration algorithm: return outfits that match filters
	// In a real implementation, this would include more sophisticated matching
	// and could suggest new outfit combinations based on wardrobe items

	return outfits, nil
}

// SubmitFeedback submits feedback for a recommendation
func (s *RecommendationServiceImpl) SubmitFeedback(recommendationID string, feedback string) error {
	if recommendationID == "" {
		return errors.New("recommendation ID is required")
	}
	if feedback == "" {
		return errors.New("feedback is required")
	}

	// Validate feedback values
	validFeedback := map[string]bool{
		"liked":    true,
		"disliked": true,
		"neutral":  true,
	}

	if !validFeedback[feedback] {
		return errors.New("feedback must be 'liked', 'disliked', or 'neutral'")
	}

	// Get the recommendation
	recommendation, err := s.recommendationRepo.GetRecommendationByID(recommendationID)
	if err != nil {
		return fmt.Errorf("recommendation not found: %w", err)
	}

	// Update feedback
	recommendation.Feedback = feedback

	return s.recommendationRepo.UpdateRecommendation(recommendation)
}

// generateOutfitRecommendations creates outfit recommendations based on user's wardrobe
// This is a helper method for future use when we want to generate new outfit combinations
func (s *RecommendationServiceImpl) generateOutfitRecommendations(userID string) ([]*domain.Outfit, error) {
	// Get user's clothing items
	items, err := s.wardrobeRepo.GetItemsByUserID(userID, map[string]interface{}{
		"isOwned": true, // Only recommend owned items
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get user wardrobe: %w", err)
	}

	if len(items) < 2 {
		return []*domain.Outfit{}, nil // Need at least 2 items to make an outfit
	}

	// Simple algorithm: create random combinations
	// In a real implementation, this would use style matching, color coordination, etc.
	var recommendations []*domain.Outfit

	// Generate a few random outfit combinations
	for i := 0; i < 3 && len(items) >= 2; i++ {
		// Pick 2-4 random items
		numItems := 2 + rand.Intn(3) // 2-4 items
		if numItems > len(items) {
			numItems = len(items)
		}

		// Shuffle items and pick first N
		shuffled := make([]*domain.ClothingItem, len(items))
		copy(shuffled, items)

		for j := range shuffled {
			k := rand.Intn(j + 1)
			shuffled[j], shuffled[k] = shuffled[k], shuffled[j]
		}

		// Create outfit from selected items
		selectedItems := shuffled[:numItems]
		itemIDs := make([]string, len(selectedItems))
		for j, item := range selectedItems {
			itemIDs[j] = item.ID
		}

		outfit := &domain.Outfit{
			UserID:        userID,
			Name:          fmt.Sprintf("Recommended Outfit %d", i+1),
			Description:   "AI-generated outfit recommendation",
			Items:         itemIDs,
			Occasion:      []string{"casual"},
			Season:        []string{"Spring", "Summer", "Fall", "Winter"},
			IsRecommended: true,
			IsFavorite:    false,
		}

		recommendations = append(recommendations, outfit)
	}

	return recommendations, nil
}
