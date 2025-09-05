package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lilo/backend/config"
	"github.com/lilo/backend/internal/handler"
	"github.com/lilo/backend/internal/repository"
	"github.com/lilo/backend/internal/service"
	"github.com/lilo/backend/pkg/middleware"
	"github.com/lilo/backend/pkg/response"
)

func main() {
	// Initialize logger
	logger := log.New(os.Stdout, "API: ", log.LstdFlags)

	// Load configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository()
	wardrobeRepo := repository.NewWardrobeRepository()
	outfitRepo := repository.NewOutfitRepository()
	recommendationRepo := repository.NewRecommendationRepository()

	// Initialize services
	userService := service.NewUserService(userRepo)
	wardrobeService := service.NewWardrobeService(wardrobeRepo)
	outfitService := service.NewOutfitService(outfitRepo)
	recommendationService := service.NewRecommendationService(recommendationRepo, wardrobeRepo, outfitRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	wardrobeHandler := handler.NewWardrobeHandler(wardrobeService)
	outfitHandler := handler.NewOutfitHandler(outfitService)
	recommendationHandler := handler.NewRecommendationHandler(recommendationService)

	// Initialize router
	router := http.NewServeMux()

	supabaseConfig := config.GetSupabaseConfig()

	// Apply middleware
	corsMiddleware := middleware.CorsMiddleware
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	authMiddleware := middleware.AuthMiddleware(userService, supabaseConfig.JWTSecret)

	// Register routes
	router.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, map[string]string{"status": "ok"})
	})

	// User routes
	router.HandleFunc("POST /api/auth/signup", userHandler.SignUp)
	router.HandleFunc("POST /api/auth/signin", userHandler.SignIn)
	router.HandleFunc("POST /api/auth/signout", userHandler.SignOut)
	router.Handle("GET /api/auth/user", authMiddleware(http.HandlerFunc(userHandler.GetUser)))
	router.Handle("PUT /api/users/profile", authMiddleware(http.HandlerFunc(userHandler.UpdateProfile)))
	router.Handle("GET /api/users/style-profile", authMiddleware(http.HandlerFunc(userHandler.GetStyleProfile)))
	router.Handle("PUT /api/users/style-profile", authMiddleware(http.HandlerFunc(userHandler.UpdateStyleProfile)))

	// Wardrobe routes
	router.Handle("GET /api/wardrobe/items", authMiddleware(http.HandlerFunc(wardrobeHandler.GetItems)))
	router.Handle("POST /api/wardrobe/items", authMiddleware(http.HandlerFunc(wardrobeHandler.AddItem)))
	router.Handle("GET /api/wardrobe/items/{id}", authMiddleware(http.HandlerFunc(wardrobeHandler.GetItem)))
	router.Handle("PUT /api/wardrobe/items/{id}", authMiddleware(http.HandlerFunc(wardrobeHandler.UpdateItem)))
	router.Handle("DELETE /api/wardrobe/items/{id}", authMiddleware(http.HandlerFunc(wardrobeHandler.DeleteItem)))
	router.Handle("GET /api/wardrobe/categories", authMiddleware(http.HandlerFunc(wardrobeHandler.GetCategories)))

	// Outfit routes
	router.Handle("GET /api/outfits", authMiddleware(http.HandlerFunc(outfitHandler.GetOutfits)))
	router.Handle("POST /api/outfits", authMiddleware(http.HandlerFunc(outfitHandler.CreateOutfit)))
	router.Handle("GET /api/outfits/{id}", authMiddleware(http.HandlerFunc(outfitHandler.GetOutfit)))
	router.Handle("PUT /api/outfits/{id}", authMiddleware(http.HandlerFunc(outfitHandler.UpdateOutfit)))
	router.Handle("DELETE /api/outfits/{id}", authMiddleware(http.HandlerFunc(outfitHandler.DeleteOutfit)))
	router.Handle("POST /api/outfits/{id}/favorite", authMiddleware(http.HandlerFunc(outfitHandler.FavoriteOutfit)))
	router.Handle("DELETE /api/outfits/{id}/favorite", authMiddleware(http.HandlerFunc(outfitHandler.UnfavoriteOutfit)))

	// Recommendation routes
	router.Handle("GET /api/recommendations/daily", authMiddleware(http.HandlerFunc(recommendationHandler.GetDaily)))
	router.Handle("GET /api/recommendations/explore", authMiddleware(http.HandlerFunc(recommendationHandler.GetExplore)))
	router.Handle("POST /api/recommendations/feedback", authMiddleware(http.HandlerFunc(recommendationHandler.SubmitFeedback)))

	// Apply global middleware
	handler := corsMiddleware(loggingMiddleware(router))

	// Create server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Printf("Server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server exited properly")
}
