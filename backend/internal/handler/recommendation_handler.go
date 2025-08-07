package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lilo/backend/internal/domain"
)

// RecommendationHandler handles recommendation-related HTTP requests
type RecommendationHandler struct {
	recommendationService domain.RecommendationService
}

// NewRecommendationHandler creates a new RecommendationHandler
func NewRecommendationHandler(recommendationService domain.RecommendationService) *RecommendationHandler {
	return &RecommendationHandler{
		recommendationService: recommendationService,
	}
}

// GetDaily returns daily outfit recommendations for the authenticated user
func (h *RecommendationHandler) GetDaily(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get daily recommendations
	recommendations, err := h.recommendationService.GetDailyRecommendations(user.ID)
	if err != nil {
		http.Error(w, "Failed to get daily recommendations", http.StatusInternalServerError)
		return
	}

	// Return recommendations
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": recommendations,
	})
}

// GetExplore returns explore recommendations for the authenticated user
func (h *RecommendationHandler) GetExplore(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Parse query parameters for filters
	filters := make(map[string]interface{})

	if occasion := r.URL.Query().Get("occasion"); occasion != "" {
		filters["occasion"] = occasion
	}
	if season := r.URL.Query().Get("season"); season != "" {
		filters["season"] = season
	}
	if isFavoriteStr := r.URL.Query().Get("isFavorite"); isFavoriteStr != "" {
		if isFavorite, err := strconv.ParseBool(isFavoriteStr); err == nil {
			filters["isFavorite"] = isFavorite
		}
	}

	// Get explore recommendations
	recommendations, err := h.recommendationService.GetExploreRecommendations(user.ID, filters)
	if err != nil {
		http.Error(w, "Failed to get explore recommendations", http.StatusInternalServerError)
		return
	}

	// Return recommendations
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": recommendations,
	})
}

// SubmitFeedback submits feedback for a recommendation
func (h *RecommendationHandler) SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	_, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var req struct {
		RecommendationID string `json:"recommendationId"`
		Feedback         string `json:"feedback"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Submit feedback
	if err := h.recommendationService.SubmitFeedback(req.RecommendationID, req.Feedback); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Feedback submitted successfully",
	})
}
