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

	passwordIsCorrect := utils.VerifyPassword(dto.Password, user.HashedPassword)

	if !passwordIsCorrect {
		return nil, models.CreateError("Invalid Password")
	}

	userGetDto := dtos.UserGetDtoFromEntity(*user)

	token, err := utils.CreateUserToken(userGetDto)

	if err != nil {
		log.Printf("%s\n", err.Error())
		return nil, models.CreateError("Cannot create token")
	}

	return token, nil
}

func (s UserService) RegisterService(registerRequest dtos.RegisterRequestDto) (*dtos.RegisterResponseDto, *models.AppError) {

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

	hashedPassword, err := utils.HashPassword(registerRequest.Password)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	registerRequest.Password = ""

	newUser := dtos.RegisterRequestDtoToEntitie(&registerRequest, role.ID, hashedPassword)

	succes, err := s.userR.CreateUser(newUser)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	if !*succes {
		return nil, models.CreateError("User could not created")
	}

	response := dtos.RegisterResponseDtoFromEntitie(newUser)

	return &response, nil
}
