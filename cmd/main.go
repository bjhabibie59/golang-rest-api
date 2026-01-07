package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	
	"go-rest-api/internal/repository"
	"go-rest-api/internal/service"
	"go-rest-api/internal/handler"
)

func main()  {
	// init layer
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// router
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")

	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8080", r))
}

