package database

import (
	"col-moda/internal/configuration"
	"col-moda/internal/domain/entities"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(configuration.AppConfiguration.DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("✅ Conexión a la base de datos exitosa")

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	// maximum number of connections.
	sqlDB.SetMaxIdleConns(10)

	// maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = Migration(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migration(db *gorm.DB) error {
	fmt.Println("🚀 Ejecutando migraciones...")

	// 🔹 Eliminar las tablas en orden inverso para evitar errores de FK
	db.Migrator().DropTable(
		&entities.Discount{}, &entities.Stripe{}, &entities.OrderDetail{}, &entities.Order{},
		&entities.Tag{}, &entities.SubCategory{}, &entities.Category{}, &entities.Product{},
		&entities.UserCompany{}, &entities.Company{}, &entities.Department{}, &entities.City{},
		&entities.User{}, &entities.Role{},
	)

	err := db.AutoMigrate(
		&entities.User{}, &entities.City{}, &entities.Department{}, &entities.Company{}, &entities.UserCompany{},
		&entities.Product{}, &entities.Category{}, &entities.SubCategory{}, &entities.Tag{}, &entities.Order{},
		&entities.OrderDetail{}, &entities.Stripe{}, &entities.Discount{},
	)
	if err != nil {
		fmt.Println("❌ Error en migraciones:", err)
		return err
	}
	fmt.Println("✅ Migraciones ejecutadas con éxito")

	err = ExecuteSQLScript(db, "cmd/database/script.sql")
	if err != nil {
		fmt.Println("❌ Error ejecutando script.sql:", err)
		return err
	}

	return nil
}

func ExecuteSQLScript(db *gorm.DB, filePath string) error {
	fmt.Println("📜 Ejecutando script SQL:", filePath)

	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error leyendo el archivo SQL: %w", err)
	}

	// Ejecutar el contenido del script.sql
	err = db.Exec(string(sqlBytes)).Error
	if err != nil {
		return fmt.Errorf("error ejecutando el script SQL: %w", err)
	}

	return nil
}
