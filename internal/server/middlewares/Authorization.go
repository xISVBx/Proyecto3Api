package middleware

import (
	"col-moda/internal/utils/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 🔹 Depuración: Imprimir headers
		fmt.Println("Headers recibidos:", c.Request.Header)

		// Obtener el header Authorization
		authHeader := c.GetHeader("Authorization")

		// Validar formato "Bearer <token>"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado o inválido"})
			c.Abort()
			return
		}

		// Extraer el token real
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verificar el token llamando a `utils.VerifyToken`
		claims, err := jwt.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Almacenar los claims en el contexto de Gin
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next() // Continuar con la petición
	}
}