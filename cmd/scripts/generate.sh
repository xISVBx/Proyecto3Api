#!/bin/bash

# region Verificar que se haya pasado un nombre
if [ -z "$1" ]; then
    echo "❌ Debes proporcionar un nombre. Ejemplo: ./generate.sh City"
    exit 1
fi

NAME=$1
NAME_LOWER=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')

#endregion

# region Obtener el nombre del módulo desde go.mod
GO_MOD_PATH="go.mod"
if [ ! -f "$GO_MOD_PATH" ]; then
    echo "❌ No se encontró go.mod en el directorio esperado."
    exit 1
fi
MODULE_NAME=$(awk '{print $2; exit}' "$GO_MOD_PATH")

echo "🔄 Creando estructura para '$NAME'..." 

#endregion

# region Crear el repositorio

mkdir -p internal/infraestructure/persistence/repositories
cat > internal/infraestructure/persistence/repositories/${NAME_LOWER}.repository.go <<EOL
package repositories

import (
    "gorm.io/gorm"
)

type ${NAME}Repository struct {
    db *gorm.DB
}

func New${NAME}Repository(db *gorm.DB) *${NAME}Repository {
    return &${NAME}Repository{db: db}
}
EOL
echo "✅ Repositorio creado: internal/infraestructure/persistence/repositories/${NAME_LOWER}.repository.go"

#endregion

#region Crear el servicio

mkdir -p internal/services
cat > internal/services/${NAME_LOWER}.service.go <<EOL
package services

import (
    "$MODULE_NAME/internal/infraestructure/persistence/repositories"
)

type ${NAME}Service struct {
    uow *repositories.UnitOfWork
}

func New${NAME}Service(uow *repositories.UnitOfWork) *${NAME}Service {
    return &${NAME}Service{uow: uow}
}
EOL
echo "✅ Servicio creado: internal/services/${NAME_LOWER}.service.go"

#endregion

#region Crear el controlador

mkdir -p internal/controllers
cat > internal/controllers/${NAME_LOWER}.controller.go <<EOL
package controllers

import (
    "$MODULE_NAME/internal/services"
    //"github.com/gin-gonic/gin"
    //"net/http"
)

type ${NAME}Controller struct {
    service *services.${NAME}Service
}

func New${NAME}Controller(service *services.${NAME}Service) *${NAME}Controller {
    return &${NAME}Controller{service: service}
}
EOL
echo "✅ Controlador creado: internal/controllers/${NAME_LOWER}.controller.go"

#endregion

#region Crear la ruta

mkdir -p internal/server/v1/routes
cat > internal/server/v1/routes/${NAME_LOWER}.route.go <<EOL
package v1_routes

import (
    "$MODULE_NAME/internal/controllers"
    "github.com/gin-gonic/gin"
)

func ${NAME}Routes(r *gin.RouterGroup, c controllers.${NAME}Controller) {
    // Definir las rutas aquí
}
EOL

echo "✅ Ruta creada: internal/routes/v1/${NAME_LOWER}.route.go"

#endregion

# region Funcion Insertar
insert_if_not_exists() {
    local FILE="$1"
    local SEARCH="$2"
    local INSERT="$3"

    # Verifica si la línea ya existe en el archivo
    if ! grep -qF -- "$INSERT" "$FILE"; then
        sed -i.bak "/$SEARCH/a \\
$INSERT\\
" "$FILE"
    fi

    # Elimina el archivo de respaldo generado por sed
    rm -f "$FILE.bak"
}
#endregion

# region Modificar UnitOfWork
UOW_FILE="internal/infraestructure/persistence/repositories/uow.go"
if [ -f "$UOW_FILE" ]; then
    echo "🔄 Modificando UnitOfWork..."
    insert_if_not_exists "$UOW_FILE" "type UnitOfWork struct {" "    ${NAME}Repo *${NAME}Repository"
    insert_if_not_exists "$UOW_FILE" "return &UnitOfWork{" "        ${NAME}Repo: New${NAME}Repository(tx),"
    echo "✅ UnitOfWork actualizado"
else
    echo "⚠️ No se encontró unit_of_work.go, no se pudo modificar."
fi

#endregion

#region Modificar servicios

SERVICES_FILE="internal/services/services.go"
if [ -f "$SERVICES_FILE" ]; then
    echo "🔄 Modificando services.go..."
    insert_if_not_exists "$SERVICES_FILE" "type Services struct {" "    ${NAME}Service *${NAME}Service"
    insert_if_not_exists "$SERVICES_FILE" "return &Services{" "        ${NAME}Service: New${NAME}Service(repos.UnitOfWork),"
    echo "✅ Services actualizado"
else
    echo "⚠️ No se encontró services.go, no se pudo modificar."
fi

#endregion

#region Modificar controladores

CONTROLLERS_FILE="internal/controllers/controllers.go"
if [ -f "$CONTROLLERS_FILE" ]; then
    echo "🔄 Modificando controllers.go..."
    
    # Agregar el nuevo controlador en la estructura Controllers
    insert_if_not_exists "$CONTROLLERS_FILE" "type Controllers struct {" "    ${NAME}Controller *${NAME}Controller"

    # Agregar la inicialización del nuevo controlador en InitControllers
    insert_if_not_exists "$CONTROLLERS_FILE" "return &Controllers{" "        ${NAME}Controller: New${NAME}Controller(services.${NAME}Service),"

    echo "✅ Controllers actualizado"
else
    echo "⚠️ No se encontró controllers.go, no se pudo modificar."
fi

#endregion

echo "🎉 Todo listo!"
