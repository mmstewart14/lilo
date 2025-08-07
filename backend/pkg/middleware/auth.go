package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lilo/backend/internal/domain"
)

// SupabaseJWTClaims represents the claims in a Supabase JWT token
type SupabaseJWTClaims struct {
	jwt.RegisteredClaims
	Email        string                 `json:"email"`
	Sub          string                 `json:"sub"` // Subject (user ID)
	Role         string                 `json:"role"`
	AppMetadata  map[string]interface{} `json:"app_metadata"`
	UserMetadata map[string]interface{} `json:"user_metadata"`
}

// AuthMiddleware creates a middleware that validates JWT tokens from Supabase
func AuthMiddleware(userService domain.UserService, jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			// Check if the header has the Bearer prefix
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
				return
			}

			// Extract the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				http.Error(w, "Token cannot be empty", http.StatusUnauthorized)
				return
			}

			// Parse and validate the token
			token, err := jwt.ParseWithClaims(tokenString, &SupabaseJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Validate the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				// Return the secret key for validation
				return []byte(jwtSecret), nil
			})

			if err != nil {
				http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
				return
			}

			// Check if the token is valid
			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Extract claims
			claims, ok := token.Claims.(*SupabaseJWTClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Check if token is expired
			if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}

			// Get user from database using the Supabase user ID
			user, err := userService.GetUserBySupabaseID(claims.Sub)
			if err != nil {
				// If user doesn't exist in our database yet, create a new user
				if errors.Is(err, domain.ErrUserNotFound) {
					// Create a new user in our database
					user = &domain.User{
						SupabaseID: claims.Sub,
						Email:      claims.Email,
						Name:       getStringFromMap(claims.UserMetadata, "name"),
						Picture:    getStringFromMap(claims.UserMetadata, "avatar_url"),
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}

					err = userService.CreateUserFromSupabase(user)
					if err != nil {
						http.Error(w, "Failed to create user", http.StatusInternalServerError)
						return
					}
				} else {
					http.Error(w, "User not found", http.StatusUnauthorized)
					return
				}
			}

			// Add the user to the request context
			ctx := context.WithValue(r.Context(), domain.ContextKeyUser, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Helper function to extract string values from a map
func getStringFromMap(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return ""
}
