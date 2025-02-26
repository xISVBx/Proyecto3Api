package dtos

import "col-moda/internal/domain/entities"

type RegisterResponse struct {
	Name     string
	LastName string
	Email    string
}

func RegisterResponseFromEntitie(user entities.User) RegisterResponse {
	return RegisterResponse{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
}
