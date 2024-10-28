package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"go-auth-jwt/config"
	"go-auth-jwt/internal/database"
	"go-auth-jwt/internal/handlers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize config
	cfg := config.New()

	// Initialize database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New())

	// Routes
	api := app.Group("/api")
	auth := api.Group("/auth")

	// Auth routes
	auth.Post("/signup", handlers.Signup(db))
	auth.Post("/login", handlers.Login(db, cfg))

	// Protected routes
	api.Get("/protected", handlers.AuthMiddleware(cfg), handlers.Protected())

	// Serve static files
	app.Static("/", "./web/static")

	// Serve HTML templates
	app.Get("/", handlers.ServeHTML("index.html"))
	app.Get("/login", handlers.ServeHTML("login.html"))
	app.Get("/signup", handlers.ServeHTML("signup.html"))
	app.Get("/authenticated", handlers.AuthMiddleware(cfg), handlers.ServeHTML("authenticated.html"))
	app.Get("/unauthenticated", handlers.ServeHTML("unauthenticated.html"))

	// Start server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
