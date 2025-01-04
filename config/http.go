package config

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// LoadHTTP Starts And Manages The Lifecycle Of An HTTP Server
func LoadHTTP(ctx context.Context, container *Container, router http.Handler) {
	server := &http.Server{
		Addr:    ":" + container.Environments.PORT,
		Handler: router,
	}

	// Goroutine To Start The Server
	go func() {
		container.Logger.Infof("Server started on port %s\n", container.Environments.PORT)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			container.Logger.Errorf("Error starting server: %s\n", err)
		}
	}()

	// Channel To Capture Signals From The Operating System
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Wait For An Interrupt Signal
	<-sig
	container.Logger.Warn("Shutting down server...")

	// Context For Shutdown With Timeout
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Shutting Down The Server
	if err := server.Shutdown(shutdownCtx); err != nil {
		container.Logger.Errorf("Server forced to shutdown: %v\n", err)
	}

	container.Logger.Warn("Server exited gracefully")
}

// Loads Route Adapters
func LoadRouter(ctx context.Context) *gin.Engine {
	// Gin Run Mode Configuration
	envMode := ctx.Value(ginModeKey).(ginMode)

	gin.SetMode(string(envMode))

	// Initialize The Router
	router := gin.New()

	// CORS Configuration
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	// Status Check Route
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Route Group
	//apiV1Router := router.Group("/v1")

	// Load Routes By Internal Adapters

	return router
}
