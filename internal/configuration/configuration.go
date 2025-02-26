package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port    string
	AppEnv  string
	DSN     string
	Jwt     []byte
	Version string
}

var AppConfiguration *Configuration

func LoadConfiguraion() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}
	if AppConfiguration == nil {
		AppConfiguration = &Configuration{
			Port:    os.Getenv("PORT"),
			AppEnv:  os.Getenv("APP_ENV"),
			DSN:     os.Getenv("DSN"),
			Jwt:     []byte(os.Getenv("JWT_SECRET")),
			Version: os.Getenv("VERSION"),
		}
	}
	log.Println("🔹 Configuración cargada:")
	log.Println("PORT:", AppConfiguration.Port)
	log.Println("APP_ENV:", AppConfiguration.AppEnv)
	log.Println("DSN:", AppConfiguration.DSN)
	log.Println("JWT_SECRET:", string(AppConfiguration.Jwt))
	log.Println("VERSION:", AppConfiguration.Version)
}
