package handler

import (
	"encoding/json"
	"net/http"

	"go-rest-api/internal/service"
	"go-rest-api/pkg/response"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	user, err := h.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to register user", nil)
		return
	}

	response.JSON(w, http.StatusCreated, "User registered successfully", user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	response.JSON(w, http.StatusOK, "Login success", map[string]string{
		"token": token,
	})
}
