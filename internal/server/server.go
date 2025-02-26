package server

import (
	"col-moda/internal/configuration"
	"col-moda/internal/controllers"
	"col-moda/internal/database"
	"col-moda/internal/infraestructure"
	"col-moda/internal/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Server struct {
	port           int
	DB             *gorm.DB
	Controllers *controllers.Controllers
}

func NewServer(db *gorm.DB) *Server {
	
	// 🔹 Inicializar repositorios, servicios y controladores
	repos := infraestructure.InitInfraestructure(db)
	services := services.InitServices(repos)
	ctrls := controllers.InitControllers(services)

	port, _ := strconv.Atoi(configuration.AppConfiguration.Port)

	return &Server{
		port:           port,
		DB:             db,
		Controllers: ctrls,
	}
}

func NewHttpServer() *http.Server {
	db, err := database.CreateConnection()

	if err != nil {
		return nil
	}

	NewServer := NewServer(db)

	handler, err := NewServer.RegisterRoutes()

	if err != nil {
		return nil
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
