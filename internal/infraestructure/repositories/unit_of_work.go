package repositories

import (
	"gorm.io/gorm"
)

// UnitOfWork gestiona la conexión y transacción si es necesario
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

// NewUnitOfWork crea la unidad de trabajo con la conexión pero sin iniciar la transacción
func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{
		db:             db, // 🔹 Usa la conexión del pool para lecturas
		DepartmentRepo: NewDepartmentRepository(db),
		UserRepo:       NewUserRepository(db),
		RoleRepo:       NewRoleRepository(db),
		CategoryRepo:   NewCategoriesRepository(db),
		ProductRepo:    NewProductRepository(db),
		CityRepo:       NewCityRepository(db),
		CompanyRepo:    NewCompanyRepository(db),
	}
}

// Begin inicia una transacción solo si se necesita
func (u *UnitOfWork) Begin() {

	if u.tx == nil {
		u.tx = u.db.Begin() // 🔹 Solo inicia transacción si es necesario
		u.DepartmentRepo.db = u.tx
		u.UserRepo.db = u.tx
		u.RoleRepo.db = u.tx
		u.CategoryRepo.db = u.tx
		u.ProductRepo.db = u.tx
		u.CityRepo.db = u.tx
		u.CompanyRepo.db = u.tx
	}
}

// Commit confirma la transacción si se inició
func (u *UnitOfWork) Commit() error {
	if u.tx != nil {
		err := u.tx.Commit().Error
		u.tx = nil 
		return err
	}
	
	return nil
}

// Rollback revierte la transacción si se inició
func (u *UnitOfWork) Rollback() error {
	if u.tx != nil {
		err := u.tx.Rollback().Error
		u.tx = nil // 🔹 Reseteamos tx después del rollback
		return err
	}
	return nil
}

// SetError guarda un error para impedir el commit
func (u *UnitOfWork) SetError(err error) {
	if err != nil {
		u.err = err
	}
}

// GetDB devuelve la base de datos actual (transacción o conexión normal)
func (u *UnitOfWork) GetDB() *gorm.DB {
	if u.tx != nil {
		return u.tx
	}
	return u.db
}
