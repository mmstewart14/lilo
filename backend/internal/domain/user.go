package domain

import (
	"errors"
	"time"
)

// User represents a user in the system
type User struct {
	ID         string    `json:"id"`
	SupabaseID string    `json:"supabaseId"`
	Email      string    `json:"email"`
	Password   string    `json:"-"` // Password is never returned in JSON
	Name       string    `json:"name,omitempty"`
	Picture    string    `json:"picture,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// StyleProfile represents a user's style preferences
type StyleProfile struct {
	ID                 string         `json:"id"`
	UserID             string         `json:"userId"`
	PreferredStyles    []string       `json:"preferredStyles"`
	WeeklySchedule     WeeklySchedule `json:"weeklySchedule"`
	SeasonalPreferences map[string][]string `json:"seasonalPreferences"`
	ColorPreferences   []string       `json:"colorPreferences"`
	UpdatedAt          time.Time      `json:"updatedAt"`
}

// WeeklySchedule represents a user's weekly clothing needs
type WeeklySchedule struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetBySupabaseID(supabaseID string) (*User, error)
	Update(user *User) error
	Delete(id string) error
	GetStyleProfile(userID string) (*StyleProfile, error)
	SaveStyleProfile(profile *StyleProfile) error
}

// Error definitions
var (
	ErrUserNotFound = errors.New("user not found")
)

// Context key for user in request context
type contextKey string
const ContextKeyUser contextKey = "user"

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(email, password, name string) (*User, error)
	CreateUserFromSupabase(user *User) error
	GetUser(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserBySupabaseID(supabaseID string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
	Authenticate(email, password string) (*User, string, error)
	VerifyToken(token string) (*User, error)
	GetStyleProfile(userID string) (*StyleProfile, error)
	SaveStyleProfile(profile *StyleProfile) error
}