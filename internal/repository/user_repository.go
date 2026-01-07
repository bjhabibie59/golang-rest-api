package repository

import "go-rest-api/internal/model"

// interface (contract) 
type UserRepository interface {
	FindAll() ([]model.User, error)
}

// implementation 
type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	users := []model.User{
		{ID: 1, Name: "Sudais", Email: "sudais@yahoo.com"},
		{ID: 1, Name: "Novan", Email: "novan@yahoo.com"},
	}

	return users, nil
}