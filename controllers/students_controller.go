package controllers

import (
	"Control_Escolar/config"
	"Control_Escolar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func GetStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

// POST
func CreateStudent(c *gin.Context) {
	var student models.Student
	// Valida (Name, Group, Email)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := config.DB.Create(&student); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// GET by ID
func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id") // Obtiene el parámetro de la URL

	if result := config.DB.First(&student, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estudiante no encontrado"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// PUT
func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if result := config.DB.First(&student, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estudiante no encontrado"})
		return
	}

	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&student).Updates(input)
	c.JSON(http.StatusOK, student)
}

// DELETE
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if result := config.DB.First(&student, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estudiante no encontrado"})
		return
	}

	config.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Estudiante eliminado"})
}
