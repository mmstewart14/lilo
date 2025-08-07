package domain

import (
	"time"
)

// Outfit represents a collection of clothing items that form an outfit
type Outfit struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	Items        []string  `json:"items"` // IDs of clothing items
	Occasion     []string  `json:"occasion"`
	Season       []string  `json:"season"`
	ImageURL     string    `json:"imageUrl,omitempty"`
	IsRecommended bool      `json:"isRecommended"`
	IsFavorite   bool      `json:"isFavorite"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// Reflection represents user feedback on an outfit they wore
type Reflection struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userId"`
	OutfitID   string    `json:"outfitId"`
	Date       time.Time `json:"date"`
	Confidence int       `json:"confidence"` // 1-5 scale
	Comfort    int       `json:"comfort"`    // 1-5 scale
	WouldRewear bool      `json:"wouldRewear"`
	Notes      string    `json:"notes,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

// Recommendation represents an outfit recommendation for a user
type Recommendation struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	OutfitID    string    `json:"outfitId"`
	Date        time.Time `json:"date"`
	Feedback    string    `json:"feedback,omitempty"` // liked, disliked, neutral
	Reason      string    `json:"reason,omitempty"`
	StylingTips []string  `json:"stylingTips"`
	CreatedAt   time.Time `json:"createdAt"`
}

// OutfitRepository defines the interface for outfit data operations
type OutfitRepository interface {
	CreateOutfit(outfit *Outfit) error
	GetOutfitByID(id string) (*Outfit, error)
	GetOutfitsByUserID(userID string, filters map[string]interface{}) ([]*Outfit, error)
	UpdateOutfit(outfit *Outfit) error
	DeleteOutfit(id string) error
	SetFavorite(id string, favorite bool) error
	CreateReflection(reflection *Reflection) error
	GetReflectionsByUserID(userID string) ([]*Reflection, error)
}

// OutfitService defines the interface for outfit business logic
type OutfitService interface {
	CreateOutfit(outfit *Outfit) error
	GetOutfit(id string) (*Outfit, error)
	GetUserOutfits(userID string, filters map[string]interface{}) ([]*Outfit, error)
	UpdateOutfit(outfit *Outfit) error
	DeleteOutfit(id string) error
	FavoriteOutfit(id string) error
	UnfavoriteOutfit(id string) error
	SubmitReflection(reflection *Reflection) error
	GetUserReflections(userID string) ([]*Reflection, error)
}

// RecommendationRepository defines the interface for recommendation data operations
type RecommendationRepository interface {
	CreateRecommendation(recommendation *Recommendation) error
	GetRecommendationByID(id string) (*Recommendation, error)
	GetRecommendationsByUserID(userID string) ([]*Recommendation, error)
	UpdateRecommendation(recommendation *Recommendation) error
}

// RecommendationService defines the interface for recommendation business logic
type RecommendationService interface {
	GetDailyRecommendations(userID string) ([]*Outfit, error)
	GetExploreRecommendations(userID string, filters map[string]interface{}) ([]*Outfit, error)
	SubmitFeedback(recommendationID string, feedback string) error
}