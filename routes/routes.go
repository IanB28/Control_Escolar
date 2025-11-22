package routes

import (
	"Control_Escolar/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Estudiantes
		api.GET("/students", controllers.GetStudents)
		api.POST("/students", controllers.CreateStudent)
		api.GET("/students/:id", controllers.GetStudent)
		api.PUT("/students/:id", controllers.UpdateStudent)
		api.DELETE("/students/:id", controllers.DeleteStudent)

		// Materias
		api.POST("/subjects", controllers.CreateSubject)
		api.GET("/subjects/:id", controllers.GetSubject)
		api.PUT("/subjects/:id", controllers.UpdateSubject)
		api.DELETE("/subjects/:id", controllers.DeleteSubject)

		// Calificaciones
		api.POST("/grades", controllers.CreateGrade)
		api.PUT("/grades/:id", controllers.UpdateGrade)
		api.DELETE("/grades/:id", controllers.DeleteGrade)
		// Rutas especificas pedidas en PDF
		api.GET("/grades/student/:student_id", controllers.GetStudentGrades)
		api.GET("/grades/:id/student/:student_id", controllers.GetGradeByStudent)
	}
}
