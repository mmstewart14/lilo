package repository

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lilo/backend/internal/domain"
)

// InMemoryRecommendationRepository implements RecommendationRepository using in-memory storage
type InMemoryRecommendationRepository struct {
	recommendations map[string]*domain.Recommendation
	mu              sync.RWMutex
}

// NewRecommendationRepository creates a new recommendation repository
func NewRecommendationRepository() domain.RecommendationRepository {
	return &InMemoryRecommendationRepository{
		recommendations: make(map[string]*domain.Recommendation),
	}
}

// CreateRecommendation creates a new recommendation
func (r *InMemoryRecommendationRepository) CreateRecommendation(recommendation *domain.Recommendation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if recommendation.ID == "" {
		recommendation.ID = uuid.New().String()
	}
	recommendation.CreatedAt = time.Now()

	r.recommendations[recommendation.ID] = recommendation
	return nil
}

// GetRecommendationByID retrieves a recommendation by ID
func (r *InMemoryRecommendationRepository) GetRecommendationByID(id string) (*domain.Recommendation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	recommendation, exists := r.recommendations[id]
	if !exists {
		return nil, domain.ErrUserNotFound // Reusing error for simplicity
	}
	return recommendation, nil
}

// GetRecommendationsByUserID retrieves all recommendations for a user
func (r *InMemoryRecommendationRepository) GetRecommendationsByUserID(userID string) ([]*domain.Recommendation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var recommendations []*domain.Recommendation
	for _, recommendation := range r.recommendations {
		if recommendation.UserID == userID {
			recommendations = append(recommendations, recommendation)
		}
	}
	return recommendations, nil
}

// UpdateRecommendation updates an existing recommendation
func (r *InMemoryRecommendationRepository) UpdateRecommendation(recommendation *domain.Recommendation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.recommendations[recommendation.ID]; !exists {
		return domain.ErrUserNotFound // Reusing error for simplicity
	}

	r.recommendations[recommendation.ID] = recommendation
	return nil
}
