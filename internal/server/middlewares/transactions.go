package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := db.Begin() // 🔥 Se inicia una nueva transacción por request
		defer func() {
			if rec := recover(); rec != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
		}()

		// Guardamos la transacción en el contexto de la request
		c.Set("tx", tx)

		// Procesamos la solicitud
		c.Next()

		// Commit solo si no hubo errores
		if c.Writer.Status() < 400 {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}
}
