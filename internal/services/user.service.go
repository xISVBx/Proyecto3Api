package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
	"col-moda/internal/utils/crypt"
	"col-moda/internal/utils/jwt"
	"log"
)

type UserService struct {
	uow *repositories.UnitOfWork
}

func NewUserService(uow *repositories.UnitOfWork) *UserService {
	return &UserService{
		uow: uow,
	}
}

func (s UserService) LoginService(dto dtos.LoginDto) (*string, *models.AppError) {

	s.uow.Begin()

	defer s.uow.Commit()

	user, err := s.uow.UserRepo.GetUserByEmail(dto.Email)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	if user == nil {
		return nil, models.NotFound("Email not found")
	}

	passwordIsCorrect := crypt.VerifyPassword(dto.Password, user.HashedPassword)

	if !passwordIsCorrect {
		return nil, models.CreateError("Invalid Password")
	}

	userGetDto := dtos.UserGetDtoFromEntity(*user)

	token, err := jwt.CreateUserToken(userGetDto)

	if err != nil {
		log.Printf("%s\n", err.Error())
		return nil, models.CreateError("Cannot create token")
	}

	return token, nil
}
 
func (s UserService) RegisterService(registerRequest dtos.RegisterRequestDto) (*dtos.RegisterResponseDto, *models.AppError) {

	s.uow.Begin()

	defer s.uow.Commit()

	user, err := s.uow.UserRepo.GetUserByEmail(registerRequest.Email)

	if err != nil {
		s.uow.Rollback()
		return nil, models.NewServerError(err)
	}

	if user != nil {
		s.uow.Rollback()
		return nil, models.CreateError("The user is already created")
	}

	role, err := s.uow.RoleRepo.FindUserRole()

	if err != nil {
		s.uow.Rollback()
		return nil, models.NewServerError(err)
	}

	hashedPassword, err := crypt.HashPassword(registerRequest.Password)

	if err != nil {
		s.uow.Rollback()
		return nil, models.NewServerError(err)
	}

	registerRequest.Password = hashedPassword

	newUser := dtos.RegisterRequestDtoToEntitie(registerRequest, role.ID)

	registerRequest.Password = ""

	succes, err := s.uow.UserRepo.CreateUser(newUser)

	if err != nil {
		s.uow.Rollback()
		return nil, models.NewServerError(err)
	}

	if !*succes {
		s.uow.Rollback()
		return nil, models.CreateError("User could not created")
	}

	response := dtos.RegisterResponseDtoFromEntitie(newUser)

	return &response, nil
}
