package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lilo/backend/internal/domain"
)

// WardrobeHandler handles wardrobe-related HTTP requests
type WardrobeHandler struct {
	wardrobeService domain.WardrobeService
}

// NewWardrobeHandler creates a new WardrobeHandler
func NewWardrobeHandler(wardrobeService domain.WardrobeService) *WardrobeHandler {
	return &WardrobeHandler{
		wardrobeService: wardrobeService,
	}
}

// GetItems returns all clothing items for the authenticated user
func (h *WardrobeHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Parse query parameters for filters
	filters := make(map[string]interface{})

	if category := r.URL.Query().Get("category"); category != "" {
		filters["category"] = category
	}
	if color := r.URL.Query().Get("color"); color != "" {
		filters["color"] = color
	}
	if season := r.URL.Query().Get("season"); season != "" {
		filters["season"] = season
	}
	if isOwnedStr := r.URL.Query().Get("isOwned"); isOwnedStr != "" {
		if isOwned, err := strconv.ParseBool(isOwnedStr); err == nil {
			filters["isOwned"] = isOwned
		}
	}

	// Get items
	items, err := h.wardrobeService.GetUserItems(user.ID, filters)
	if err != nil {
		http.Error(w, "Failed to get wardrobe items", http.StatusInternalServerError)
		return
	}

	// Return items
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": items,
	})
}

// AddItem adds a new clothing item to the user's wardrobe
func (h *WardrobeHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var item domain.ClothingItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set user ID
	item.UserID = user.ID

	// Add item
	if err := h.wardrobeService.AddItem(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return created item
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Item added successfully",
		"data":    item,
	})
}

// GetItem returns a specific clothing item by ID
func (h *WardrobeHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	// Get item ID from URL path
	itemID := r.PathValue("id")
	if itemID == "" {
		http.Error(w, "Item ID is required", http.StatusBadRequest)
		return
	}

	// Get item
	item, err := h.wardrobeService.GetItem(itemID)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	// Verify user owns the item
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok || item.UserID != user.ID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	// Return item
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": item,
	})
}

// UpdateItem updates an existing clothing item
func (h *WardrobeHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get item ID from URL path
	itemID := r.PathValue("id")
	if itemID == "" {
		http.Error(w, "Item ID is required", http.StatusBadRequest)
		return
	}

	// Parse request body
	var item domain.ClothingItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set IDs
	item.ID = itemID
	item.UserID = user.ID

	// Update item
	if err := h.wardrobeService.UpdateItem(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return updated item
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Item updated successfully",
		"data":    item,
	})
}

// DeleteItem deletes a clothing item
func (h *WardrobeHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get item ID from URL path
	itemID := r.PathValue("id")
	if itemID == "" {
		http.Error(w, "Item ID is required", http.StatusBadRequest)
		return
	}

	// Verify user owns the item before deletion
	item, err := h.wardrobeService.GetItem(itemID)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	if item.UserID != user.ID {
		http.Error(w, "Unauthorized", http.StatusForbidden)
		return
	}

	// Delete item
	if err := h.wardrobeService.DeleteItem(itemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Item deleted successfully",
	})
}

// GetCategories returns all available clothing categories
func (h *WardrobeHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	// Get categories
	categories, err := h.wardrobeService.GetCategories()
	if err != nil {
		http.Error(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}

	// Return categories
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": categories,
	})
}
