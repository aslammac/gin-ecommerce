package main

import (
	"gin-ecommerce/handlers"
	"gin-ecommerce/middleware"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		println("Error loading .env file")
	}

	r := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Use environment variable
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = false
	r.Use(cors.New(config))

	// Public routes
	r.POST("/auth/send-otp", handlers.SendOTP)
	r.POST("/auth/verify-otp", handlers.VerifyOTP)
	r.POST("/auth/signup", handlers.Signup)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/products", handlers.GetProducts)
		protected.POST("/products", handlers.CreateProduct)
		protected.GET("/orders", handlers.GetOrders)
		protected.POST("/orders", handlers.CreateOrder)
		protected.GET("/users/profile", handlers.GetUserProfile)
		protected.PUT("/users/profile", handlers.UpdateUserProfile)
		// Add more e-commerce related endpoints here
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	r.Run(":" + port)
}