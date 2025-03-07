package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
	"col-moda/internal/utils/crypt"
)

type CompanyService struct {
	uow *repositories.UnitOfWork
}

func NewCompanyService(uow *repositories.UnitOfWork) *CompanyService {
	return &CompanyService{
		uow: uow,
	}
}

// FindCompaniesByFilters busca empresas aplicando filtros opcionales
func (r *CompanyService) FindCompaniesByFilters(dto dtos.CompanyRequestDto) ([]dtos.CompanyResponseDto, *models.AppError) {
	var companies, err = r.uow.CompanyRepo.FindCompaniesByFilters(dto)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	companiesResponse := dtos.CompaniesResponseDtoFromEntities(companies)

	return companiesResponse, nil
}

func (s CompanyService) RegisterCompany(registerCompany dtos.RegisterCompanyRequestDto) (*bool, *models.AppError) {

	//Validar si el usuario ya existe

	user, err := s.uow.UserRepo.GetUserByEmail(registerCompany.Email)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	if user != nil {
		return nil, models.CreateError("The user is already created")
	}

	//Buscar el rol de compañia

	role, err := s.uow.RoleRepo.FindCompanyRole()

	if err != nil {
		return nil, models.NewServerError(err)
	}

	//Hashear la contraseña

	hashedPassword, err := crypt.HashPassword(registerCompany.Password)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	registerCompany.Password = hashedPassword

	newUser, newCompany, newUserCompany := dtos.RegisterCompanyRequestDtoToEntitie(registerCompany, role.ID)

	registerCompany.Password = ""

	//Crear el usuario

	succes, err := s.uow.UserRepo.CreateUser(newUser)

	if err != nil {
		s.uow.Rollback()
		return nil, models.NewServerError(err)
	}

	if !*succes {
		s.uow.Rollback()
		return nil, models.CreateError("User could not created")
	}

	//Crear la compañia

	succes, err = s.uow.CompanyRepo.CreateCompany(newCompany)

	if err != nil {
		s.uow.Rollback()
		return nil, models.NewServerError(err)
	}

	if !*succes {
		s.uow.Rollback()
		return nil, models.CreateError("User could not created")
	}

	//Crear Usuario Compañia

	succes, err = s.uow.CompanyRepo.CreateUserCompany(&newUserCompany)

	if err != nil {
		s.uow.Rollback()
		return nil, models.NewServerError(err)
	}

	if !*succes {
		s.uow.Rollback()
		return nil, models.CreateError("User could not created")
	}

	err = s.uow.Commit()

	if err != nil {
		return nil, models.NewServerError(err)
	}

	return succes, nil
}
