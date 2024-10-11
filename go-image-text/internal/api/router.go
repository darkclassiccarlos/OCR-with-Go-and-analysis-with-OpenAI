package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configura todas las rutas de la API y aplica middleware
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Aplicar middleware global
	router.Use(LoggerMiddleware())        // Para registrar cada solicitud
	router.Use(CORSMiddleware())          // Para manejar CORS
	router.Use(ErrorHandlingMiddleware()) // Para manejar errores globalmente

	// Rutas
	router.POST("/analyze", AnalyzeImage)

	return router
}
