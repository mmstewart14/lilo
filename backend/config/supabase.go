package config

import (
	"os"
)

// SupabaseConfig holds Supabase configuration
type SupabaseConfig struct {
	URL       string
	AnonKey   string
	JWTSecret string
}

// GetSupabaseConfig returns the Supabase configuration
func GetSupabaseConfig() *SupabaseConfig {
	return &SupabaseConfig{
		URL:       getEnvOrDefault("SUPABASE_URL", ""),
		AnonKey:   getEnvOrDefault("SUPABASE_ANON_KEY", ""),
		JWTSecret: getEnvOrDefault("SUPABASE_JWT_SECRET", ""),
	}
}

// getEnvOrDefault returns the environment variable value or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
