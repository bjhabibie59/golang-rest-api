package handler

import (
	"net/http"

	"go-rest-api/internal/service"
	"go-rest-api/pkg/response"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Failed to fetch users", nil)
		return
	}

	response.JSON(w, http.StatusOK, "Success", users)
}