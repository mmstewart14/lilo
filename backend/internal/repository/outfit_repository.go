package repository

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lilo/backend/internal/domain"
)

// InMemoryOutfitRepository implements OutfitRepository using in-memory storage
type InMemoryOutfitRepository struct {
	outfits     map[string]*domain.Outfit
	reflections map[string]*domain.Reflection
	mu          sync.RWMutex
}

// NewOutfitRepository creates a new outfit repository
func NewOutfitRepository() domain.OutfitRepository {
	return &InMemoryOutfitRepository{
		outfits:     make(map[string]*domain.Outfit),
		reflections: make(map[string]*domain.Reflection),
	}
}

// CreateOutfit creates a new outfit
func (r *InMemoryOutfitRepository) CreateOutfit(outfit *domain.Outfit) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if outfit.ID == "" {
		outfit.ID = uuid.New().String()
	}
	outfit.CreatedAt = time.Now()
	outfit.UpdatedAt = time.Now()

	r.outfits[outfit.ID] = outfit
	return nil
}

// GetOutfitByID retrieves an outfit by ID
func (r *InMemoryOutfitRepository) GetOutfitByID(id string) (*domain.Outfit, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	outfit, exists := r.outfits[id]
	if !exists {
		return nil, domain.ErrUserNotFound // Reusing error for simplicity
	}
	return outfit, nil
}

// GetOutfitsByUserID retrieves all outfits for a user with optional filters
func (r *InMemoryOutfitRepository) GetOutfitsByUserID(userID string, filters map[string]interface{}) ([]*domain.Outfit, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var outfits []*domain.Outfit
	for _, outfit := range r.outfits {
		if outfit.UserID == userID {
			// Apply filters if provided
			if r.matchesFilters(outfit, filters) {
				outfits = append(outfits, outfit)
			}
		}
	}
	return outfits, nil
}

// matchesFilters checks if an outfit matches the provided filters
func (r *InMemoryOutfitRepository) matchesFilters(outfit *domain.Outfit, filters map[string]interface{}) bool {
	if filters == nil {
		return true
	}

	if isFavorite, ok := filters["isFavorite"]; ok {
		if outfit.IsFavorite != isFavorite.(bool) {
			return false
		}
	}

	if isRecommended, ok := filters["isRecommended"]; ok {
		if outfit.IsRecommended != isRecommended.(bool) {
			return false
		}
	}

	if occasion, ok := filters["occasion"]; ok {
		occasionStr := occasion.(string)
		found := false
		for _, o := range outfit.Occasion {
			if o == occasionStr {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if season, ok := filters["season"]; ok {
		seasonStr := season.(string)
		found := false
		for _, s := range outfit.Season {
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

// UpdateOutfit updates an existing outfit
func (r *InMemoryOutfitRepository) UpdateOutfit(outfit *domain.Outfit) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.outfits[outfit.ID]; !exists {
		return domain.ErrUserNotFound // Reusing error for simplicity
	}

	outfit.UpdatedAt = time.Now()
	r.outfits[outfit.ID] = outfit
	return nil
}

// DeleteOutfit deletes an outfit by ID
func (r *InMemoryOutfitRepository) DeleteOutfit(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.outfits[id]; !exists {
		return domain.ErrUserNotFound // Reusing error for simplicity
	}

	delete(r.outfits, id)
	return nil
}

// SetFavorite sets the favorite status of an outfit
func (r *InMemoryOutfitRepository) SetFavorite(id string, favorite bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	outfit, exists := r.outfits[id]
	if !exists {
		return domain.ErrUserNotFound // Reusing error for simplicity
	}

	outfit.IsFavorite = favorite
	outfit.UpdatedAt = time.Now()
	r.outfits[id] = outfit
	return nil
}

// CreateReflection creates a new reflection
func (r *InMemoryOutfitRepository) CreateReflection(reflection *domain.Reflection) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if reflection.ID == "" {
		reflection.ID = uuid.New().String()
	}
	reflection.CreatedAt = time.Now()

	r.reflections[reflection.ID] = reflection
	return nil
}

// GetReflectionsByUserID retrieves all reflections for a user
func (r *InMemoryOutfitRepository) GetReflectionsByUserID(userID string) ([]*domain.Reflection, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var reflections []*domain.Reflection
	for _, reflection := range r.reflections {
		if reflection.UserID == userID {
			reflections = append(reflections, reflection)
		}
	}
	return reflections, nil
}
