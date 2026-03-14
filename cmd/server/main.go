package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imwaddy/url-shortner/config"
	"github.com/imwaddy/url-shortner/handler"
	"github.com/imwaddy/url-shortner/pkg/cache"
	"github.com/imwaddy/url-shortner/pkg/database"
	"github.com/imwaddy/url-shortner/pkg/logger"
	"github.com/imwaddy/url-shortner/repository"
	"github.com/imwaddy/url-shortner/service"
)

func main() {
	logger.Init()
	logger.Println("🚀 Starting URL Shortener Service...")

	// Load configuration
	cfg := config.Load()

	// Initialize Redis
	redisCache := cache.NewRedisClient(cfg.RedisAddr)

	// Initialize MySQL
	db, err := database.NewMySQL(cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	logger.Println("✅ Connected to MySQL")

	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	logger.Println("✅ Database migrations completed")

	// Initialize layers
	repo := repository.NewURLRepository(db, redisCache)
	svc := service.NewURLService(repo)
	h := handler.NewURLHandler(svc, cfg.BaseURL)

	// Setup Gin router
	r := gin.Default()
	h.RegisterRoutes(r)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		logger.Printf("🌐 Server listening on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("🛑 Shutting down server...")

	// Graceful shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Close database connection
	if err := db.Close(); err != nil {
		logger.Errorf("Error closing database: %v", err)
	}

	// Close Redis connection
	if err := redisCache.Close(); err != nil {
		logger.Errorf("Error closing Redis: %v", err)
	}

	logger.Println("✅ Server exited gracefully")
}
