package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
	"col-moda/internal/utils"
	"log"
)

type UserService struct {
	userR *repositories.UserRepository
	roleR *repositories.RoleRepository
}

func NewUserService(userR *repositories.UserRepository, roleR *repositories.RoleRepository) *UserService {
	return &UserService{
		userR: userR,
		roleR: roleR,
	}
}

func (s UserService) LoginService(dto dtos.LoginDto) (*string, *models.AppError) {

	user, err := s.userR.GetUserByEmail(dto.Email)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	if user == nil {
		return nil, models.NotFound("Email not found")
	}

	userGetDto := dtos.UserGetDtoFromEntity(*user)

	token, err := utils.CreateUserToken(userGetDto)

	if err != nil {
		log.Printf("%s\n", err.Error())
		return nil, models.NotFound("Cannot create token")
	}

	return token, nil
}

func (s UserService) RegisterService(registerRequest dtos.RegisterRequest) (*dtos.RegisterResponse, *models.AppError) {

	user, err := s.userR.GetUserByEmail(registerRequest.Email)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	if user != nil {
		return nil, models.CreateError("The user is already created")
	}

	role, err := s.roleR.FindUserRole()

	if err != nil {
		return nil, models.NewServerError(err)
	}

	newUser := dtos.RegisterRequestToUserEntity(&registerRequest, role.ID)

	succes, err := s.userR.CreateUser(newUser)

	if err != nil {
		models.NewServerError(err)
	}

	if !*succes {
		models.CreateError("User could not created")
	}

	response := dtos.RegisterResponseFromEntitie(newUser)

	return &response, nil

}
