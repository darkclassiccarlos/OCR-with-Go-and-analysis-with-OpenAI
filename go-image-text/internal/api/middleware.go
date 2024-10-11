package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware para logging de cada solicitud
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Registrar la solicitud entrante
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Continuar con la siguiente función/middleware
		c.Next()

		// Al final de la solicitud, calcular el tiempo de procesamiento y registrar los detalles
		duration := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("%s %s %d %v", method, path, statusCode, duration)
	}
}

// Middleware para manejo de CORS (permite peticiones desde cualquier origen)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// Middleware para manejo de errores
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Atrapar cualquier error que ocurra durante la solicitud
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Ocurrió un error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Ha ocurrido un error interno",
				})
				c.Abort()
			}
		}()
		// Continuar con la siguiente función/middleware
		c.Next()
	}
}
