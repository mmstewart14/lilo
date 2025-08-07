package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lilo/backend/internal/domain"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userService domain.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// SignUp handles user registration from Supabase
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req struct {
		SupabaseID string `json:"supabaseId"`
		Email      string `json:"email"`
		Name       string `json:"name"`
		Picture    string `json:"picture"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create user in our database
	user := &domain.User{
		SupabaseID: req.SupabaseID,
		Email:      req.Email,
		Name:       req.Name,
		Picture:    req.Picture,
	}

	if err := h.userService.CreateUserFromSupabase(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"data": user,
	})
}

// SignIn handles user authentication (not needed for Supabase Auth)
func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// Supabase handles authentication, so this is just a placeholder
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Authentication is handled by Supabase",
	})
}

// SignOut handles user sign out (not needed for Supabase Auth)
func (h *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	// Supabase handles sign out, so this is just a placeholder
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Sign out is handled by Supabase",
	})
}

// GetUser returns the current authenticated user
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Get user from context (set by auth middleware)
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Return user data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": user,
	})
}

// UpdateProfile updates the user's profile
func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var req struct {
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update user data
	user.Name = req.Name
	user.Picture = req.Picture

	if err := h.userService.UpdateUser(user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Return updated user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Profile updated successfully",
		"data": user,
	})
}

// GetStyleProfile returns the user's style profile
func (h *UserHandler) GetStyleProfile(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Get style profile
	profile, err := h.userService.GetStyleProfile(user.ID)
	if err != nil {
		http.Error(w, "Failed to get style profile", http.StatusInternalServerError)
		return
	}

	// Return style profile
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": profile,
	})
}

// UpdateStyleProfile updates the user's style profile
func (h *UserHandler) UpdateStyleProfile(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user, ok := r.Context().Value(domain.ContextKeyUser).(*domain.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var profile domain.StyleProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set user ID
	profile.UserID = user.ID

	// Save style profile
	if err := h.userService.SaveStyleProfile(&profile); err != nil {
		http.Error(w, "Failed to save style profile", http.StatusInternalServerError)
		return
	}

	// Return updated profile
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Style profile updated successfully",
		"data": profile,
	})
}