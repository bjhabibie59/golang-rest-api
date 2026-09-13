package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go-rest-api/internal/config"
	"go-rest-api/internal/handler"
	"go-rest-api/internal/middleware"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/service"
)

func main() {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Initialize Database Connection & Auto Migration / Seeding
	db := config.ConnectDB()

	// init layer
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	// router
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Protected routes (using AuthMiddleware)
	r.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(userHandler.GetUsers))).Methods("GET")

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
