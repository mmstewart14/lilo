package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lilo/backend/internal/domain"
)

// OutfitHandler handles outfit-related HTTP requests
type OutfitHandler struct {
	outfitService domain.OutfitService
}

// NewOutfitHandler creates a new OutfitHandler
func NewOutfitHandler(outfitService domain.OutfitService) *OutfitHandler {
	return &OutfitHandler{
		outfitService: outfitService,
	}
}

// GetOutfits returns all outfits for the authenticated user
func (h *OutfitHandler) GetOutfits(w http.ResponseWriter, r *http.Request) {
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
	if isRecommendedStr := r.URL.Query().Get("isRecommended"); isRecommendedStr != "" {
		if isRecommended, err := strconv.ParseBool(isRecommendedStr); err == nil {
			filters["isRecommended"] = isRecommended
		}
	}

	// Get outfits
	outfits, err := h.outfitService.GetUserOutfits(user.ID, filters)
	if err != nil {
		http.Error(w, "Failed to get outfits", http.StatusInternalServerError)
		return
	}

	// Return outfits
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": outfits,
	})
}

// CreateOutfit creates a new outfit
func (h *OutfitHandler) CreateOutfit(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var outfit domain.Outfit
	if err := json.NewDecoder(r.Body).Decode(&outfit); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set user ID
	outfit.UserID = user.ID

	// Create outfit
	if err := h.outfitService.CreateOutfit(&outfit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return created outfit
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Outfit created successfully",
		"data":    outfit,
	})
}

// GetOutfit returns a specific outfit by ID
func (h *OutfitHandler) GetOutfit(w http.ResponseWriter, r *http.Request) {
	// Get outfit ID from URL path
	outfitID := r.PathValue("id")
	if outfitID == "" {
		http.Error(w, "Outfit ID is required", http.StatusBadRequest)
		return
	}

	// Get outfit
	outfit, err := h.outfitService.GetOutfit(outfitID)
	if err != nil {
		http.Error(w, "Outfit not found", http.StatusNotFound)
		return
	}

	// Verify user owns the outfit
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok || outfit.UserID != user.ID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	// Return outfit
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": outfit,
	})
}

// UpdateOutfit updates an existing outfit
func (h *OutfitHandler) UpdateOutfit(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get outfit ID from URL path
	outfitID := r.PathValue("id")
	if outfitID == "" {
		http.Error(w, "Outfit ID is required", http.StatusBadRequest)
		return
	}

	// Parse request body
	var outfit domain.Outfit
	if err := json.NewDecoder(r.Body).Decode(&outfit); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set IDs
	outfit.ID = outfitID
	outfit.UserID = user.ID

	// Update outfit
	if err := h.outfitService.UpdateOutfit(&outfit); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return updated outfit
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Outfit updated successfully",
		"data":    outfit,
	})
}

// DeleteOutfit deletes an outfit
func (h *OutfitHandler) DeleteOutfit(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get outfit ID from URL path
	outfitID := r.PathValue("id")
	if outfitID == "" {
		http.Error(w, "Outfit ID is required", http.StatusBadRequest)
		return
	}

	// Verify user owns the outfit before deletion
	outfit, err := h.outfitService.GetOutfit(outfitID)
	if err != nil {
		http.Error(w, "Outfit not found", http.StatusNotFound)
		return
	}

	if outfit.UserID != user.ID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	// Delete outfit
	if err := h.outfitService.DeleteOutfit(outfitID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Outfit deleted successfully",
	})
}

// FavoriteOutfit marks an outfit as favorite
func (h *OutfitHandler) FavoriteOutfit(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get outfit ID from URL path
	outfitID := r.PathValue("id")
	if outfitID == "" {
		http.Error(w, "Outfit ID is required", http.StatusBadRequest)
		return
	}

	// Verify user owns the outfit
	outfit, err := h.outfitService.GetOutfit(outfitID)
	if err != nil {
		http.Error(w, "Outfit not found", http.StatusNotFound)
		return
	}

	if outfit.UserID != user.ID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	// Favorite outfit
	if err := h.outfitService.FavoriteOutfit(outfitID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Outfit favorited successfully",
	})
}

// UnfavoriteOutfit removes favorite status from an outfit
func (h *OutfitHandler) UnfavoriteOutfit(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get outfit ID from URL path
	outfitID := r.PathValue("id")
	if outfitID == "" {
		http.Error(w, "Outfit ID is required", http.StatusBadRequest)
		return
	}

	// Verify user owns the outfit
	outfit, err := h.outfitService.GetOutfit(outfitID)
	if err != nil {
		http.Error(w, "Outfit not found", http.StatusNotFound)
		return
	}

	if outfit.UserID != user.ID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	// Unfavorite outfit
	if err := h.outfitService.UnfavoriteOutfit(outfitID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Outfit unfavorited successfully",
	})
}
