package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
		URL:       getEnvVar("SUPABASE_URL"),
		AnonKey:   getEnvVar("SUPABASE_ANON_KEY"),
		JWTSecret: getEnvVar("SUPABASE_JWT_SECRET"),
	}
}

// Getenv returns the environment variable value or a default value
func getEnvVar(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}
