package controllers

import (
	"Control_Escolar/config"
	"Control_Escolar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST
func CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&subject)
	c.JSON(http.StatusCreated, subject)
}

// GET
func GetSubject(c *gin.Context) {
	var subject models.Subject
	id := c.Param("id")
	if result := config.DB.First(&subject, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
		return
	}
	c.JSON(http.StatusOK, subject)
}

// PUT
func UpdateSubject(c *gin.Context) {
	var subject models.Subject
	id := c.Param("id")
	if result := config.DB.First(&subject, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
		return
	}
	var input models.Subject
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&subject).Updates(input)
	c.JSON(http.StatusOK, subject)
}

// DELETE
func DeleteSubject(c *gin.Context) {
	var subject models.Subject
	id := c.Param("id")
	if result := config.DB.First(&subject, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
		return
	}
	config.DB.Delete(&subject)
	c.JSON(http.StatusOK, gin.H{"message": "Materia eliminada"})
}
