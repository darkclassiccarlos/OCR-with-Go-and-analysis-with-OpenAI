package main

import (
	"go-image-text/internal/api"
)

func main() {
	router := api.SetupRouter() // Configura el router
	router.Run(":8080")         // Levanta el servidor
}
