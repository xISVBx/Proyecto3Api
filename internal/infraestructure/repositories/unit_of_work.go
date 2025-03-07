package repositories

import (
	"gorm.io/gorm"
)

// UnitOfWork gestiona una única transacción y sus repositorios.
type UnitOfWork struct {
	db             *gorm.DB
	tx             *gorm.DB
	err            error
	DepartmentRepo *DepartmentRepository
	UserRepo       *UserRepository
	RoleRepo       *RoleRepository
	CategoryRepo   *CategoryRepository
	ProductRepo    *ProductRepository
	CityRepo       *CityRepository
	CompanyRepo    *CompanyRepository
}

// NewUnitOfWork crea una nueva instancia con una transacción activa y repositorios.
func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	tx := db.Begin() // Se inicia la transacción

	return &UnitOfWork{
		db:             db,
		tx:             tx,
		DepartmentRepo: NewDepartmentRepository(tx),  
		UserRepo:       NewUserRepository(tx),
		RoleRepo:       NewRoleRepository(tx),
		CategoryRepo:   NewCategoriesRepository(tx),
		ProductRepo:    NewProductRepository(tx),
		CityRepo:       NewCityRepository(tx),
		CompanyRepo:    NewCompanyRepository(tx),
	}
}


// Commit confirma la transacción si no hubo errores.
func (u *UnitOfWork) Commit() error {
	if u.err != nil {
		return u.err
	}
	return u.tx.Commit().Error
}

// Rollback revierte la transacción.
func (u *UnitOfWork) Rollback() error {
	return u.tx.Rollback().Error
}

// SetError guarda un error para impedir el commit.
func (u *UnitOfWork) SetError(err error) {
	if err != nil {
		u.err = err
	}
}
