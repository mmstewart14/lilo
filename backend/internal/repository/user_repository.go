package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lilo/backend/internal/domain"
)

// InMemoryUserRepository implements UserRepository using in-memory storage
type InMemoryUserRepository struct {
	users         map[string]*domain.User
	styleProfiles map[string]*domain.StyleProfile
	mu            sync.RWMutex
}

// NewUserRepository creates a new user repository
func NewUserRepository() domain.UserRepository {
	return &InMemoryUserRepository{
		users:         make(map[string]*domain.User),
		styleProfiles: make(map[string]*domain.StyleProfile),
	}
}

// Create creates a new user
func (r *InMemoryUserRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by ID
func (r *InMemoryUserRepository) GetByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

// GetByEmail retrieves a user by email
func (r *InMemoryUserRepository) GetByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

// GetBySupabaseID retrieves a user by Supabase ID
func (r *InMemoryUserRepository) GetBySupabaseID(supabaseID string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.SupabaseID == supabaseID {
			return user, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

// Update updates an existing user
func (r *InMemoryUserRepository) Update(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return domain.ErrUserNotFound
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return nil
}

// Delete deletes a user by ID
func (r *InMemoryUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return domain.ErrUserNotFound
	}

	delete(r.users, id)
	delete(r.styleProfiles, id) // Also delete style profile
	return nil
}

// GetStyleProfile retrieves a user's style profile
func (r *InMemoryUserRepository) GetStyleProfile(userID string) (*domain.StyleProfile, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	profile, exists := r.styleProfiles[userID]
	if !exists {
		return nil, fmt.Errorf("style profile not found for user %s", userID)
	}
	return profile, nil
}

// SaveStyleProfile saves a user's style profile
func (r *InMemoryUserRepository) SaveStyleProfile(profile *domain.StyleProfile) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if profile.ID == "" {
		profile.ID = uuid.New().String()
	}
	profile.UpdatedAt = time.Now()

	r.styleProfiles[profile.UserID] = profile
	return nil
}
