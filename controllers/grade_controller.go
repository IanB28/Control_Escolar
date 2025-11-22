package controllers

import (
	"Control_Escolar/config"
	"Control_Escolar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST
func CreateGrade(c *gin.Context) {
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Valida la existencia del estudiante
	var student models.Student
	if result := config.DB.First(&student, grade.StudentID); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID de estudiante no existe"})
		return
	}

	// Valida la existencia de la materia
	var subject models.Subject
	if result := config.DB.First(&subject, grade.SubjectID); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID de materia no existe"})
		return
	}

	config.DB.Create(&grade)
	c.JSON(http.StatusCreated, grade)
}

// PUT
func UpdateGrade(c *gin.Context) {
	var grade models.Grade
	id := c.Param("id")
	if result := config.DB.First(&grade, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada"})
		return
	}
	var input models.Grade
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&grade).Updates(input)
	c.JSON(http.StatusOK, grade)
}

// DELETE
func DeleteGrade(c *gin.Context) {
	var grade models.Grade
	id := c.Param("id")
	if result := config.DB.First(&grade, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada"})
		return
	}
	config.DB.Delete(&grade)
	c.JSON(http.StatusOK, gin.H{"message": "Calificación eliminada"})
}

// GET /api/grades/student/:student_id
func GetStudentGrades(c *gin.Context) {
	var grades []models.Grade
	studentID := c.Param("student_id")
	config.DB.Where("student_id = ?", studentID).Preload("Subject").Preload("Student").Find(&grades)
	c.JSON(http.StatusOK, grades)
}

// GET /api/grades/:id/student/:student_id
func GetGradeByStudent(c *gin.Context) {
	var grade models.Grade
	gradeID := c.Param("id")
	studentID := c.Param("student_id")

	if result := config.DB.Where("id = ? AND student_id = ?", gradeID, studentID).Preload("Subject").First(&grade); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada para este estudiante"})
		return
	}
	c.JSON(http.StatusOK, grade)
}
