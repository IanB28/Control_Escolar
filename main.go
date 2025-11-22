package main

import (
	"Control_Escolar/config"
	"Control_Escolar/models"
	"Control_Escolar/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Student{}, &models.Subject{}, &models.Grade{})
	fmt.Println("Tablas migradas correctamente")

	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("SERVER_PORT")

	fmt.Printf("Servidor corriendo en el puerto %s\n", port)
	r.Run(":" + port)
}
