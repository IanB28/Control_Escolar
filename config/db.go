package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Datos de Acceso a la Base de Datos
	dsn := "host=localhost user=postgres password=Admin dbname=control_escolar port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	DB = database
	fmt.Println("Conexion a Base de Datos Exitosa")
}
