# Variables para evitar repetir rutas
MAIN_FILE=cmd/api/main.go  # Archivo principal de la API
SWAGGER_DIR=docs/swagger   # Carpeta donde se genera Swagger

# Construir la aplicación
all: swag build test

build: swag
	@echo "Building..."
	@go build -o tmp/main $(MAIN_FILE)

# Ejecutar la aplicación
run: swag
	@go run $(MAIN_FILE)

# Generar Swagger (solo antes de ejecutar la app)
swag:
	@echo "Generating Swagger docs..."
	@if [ ! -d "$(SWAGGER_DIR)" ]; then mkdir -p $(SWAGGER_DIR); fi
	@swag init -g $(MAIN_FILE) -o $(SWAGGER_DIR)

# Ejecutar pruebas
test:
	@echo "Testing..."
	@go test ./... -v

# Limpiar archivos generados
clean:
	@echo "Cleaning..."
	@rm -f tmp/main

# Live Reload (sin regenerar Swagger en cada cambio)
watch:
	@echo "Generating Swagger docs (one-time before watch)..."
	@make swag
	@if command -v air > /dev/null; then \
		echo "Watching..."; \
		air -c .air.toml; \
	else \
		echo "Installing air..."; \
		go install github.com/air-verse/air@latest; \
		air -c .air.toml; \
	fi

# Alias para generar estructura de repositorio, servicio y controlador
generate:
	bash cmd/scripts/generate.sh $(name)



.PHONY: all build run test clean watch swag
