package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/lilo/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserServiceImpl implements UserService
type UserServiceImpl struct {
	userRepo domain.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user with hashed password
func (s *UserServiceImpl) CreateUser(email, password, name string) (*domain.User, error) {
	// Check if user already exists
	if _, err := s.userRepo.GetByEmail(email); err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &domain.User{
		Email:     email,
		Password:  string(hashedPassword),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// CreateUserFromSupabase creates a user from Supabase authentication
func (s *UserServiceImpl) CreateUserFromSupabase(user *domain.User) error {
	// Check if user already exists by Supabase ID
	if _, err := s.userRepo.GetBySupabaseID(user.SupabaseID); err == nil {
		return errors.New("user already exists")
	}

	return s.userRepo.Create(user)
}

// GetUser retrieves a user by ID
func (s *UserServiceImpl) GetUser(id string) (*domain.User, error) {
	return s.userRepo.GetByID(id)
}

// GetUserByEmail retrieves a user by email
func (s *UserServiceImpl) GetUserByEmail(email string) (*domain.User, error) {
	return s.userRepo.GetByEmail(email)
}

// GetUserBySupabaseID retrieves a user by Supabase ID
func (s *UserServiceImpl) GetUserBySupabaseID(supabaseID string) (*domain.User, error) {
	return s.userRepo.GetBySupabaseID(supabaseID)
}

// UpdateUser updates an existing user
func (s *UserServiceImpl) UpdateUser(user *domain.User) error {
	return s.userRepo.Update(user)
}

// DeleteUser deletes a user by ID
func (s *UserServiceImpl) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

// Authenticate authenticates a user with email and password
func (s *UserServiceImpl) Authenticate(email, password string) (*domain.User, string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate a simple token (in production, use JWT)
	token, err := s.generateToken()
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return user, token, nil
}

// VerifyToken verifies a token and returns the associated user
func (s *UserServiceImpl) VerifyToken(token string) (*domain.User, error) {
	// This is a simplified implementation
	// In production, you would verify JWT tokens properly
	return nil, errors.New("token verification not implemented")
}

// GetStyleProfile retrieves a user's style profile
func (s *UserServiceImpl) GetStyleProfile(userID string) (*domain.StyleProfile, error) {
	return s.userRepo.GetStyleProfile(userID)
}

// SaveStyleProfile saves a user's style profile
func (s *UserServiceImpl) SaveStyleProfile(profile *domain.StyleProfile) error {
	return s.userRepo.SaveStyleProfile(profile)
}

// generateToken generates a simple random token
func (s *UserServiceImpl) generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
