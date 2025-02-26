package dtos

import "col-moda/internal/domain/entities"

type UserGetDto struct {
	Name  string
	Email string
}

func UserGetDtoFromEntity(user entities.User) UserGetDto {
	return UserGetDto{
		Name:  user.Name,
		Email: user.Email,
	}
}

