package models

type Grade struct {
	ID        uint    `gorm:"primaryKey" json:"grade_id"`
	StudentID uint    `json:"student_id" binding:"required"`
	SubjectID uint    `json:"subject_id" binding:"required"`
	Grade     float64 `json:"grade" binding:"required,gte=0,lte=10"` // Validación 0-10

	// FK
	Student Student `gorm:"foreignKey:StudentID" binding:"-" json:"-"`
	Subject Subject `gorm:"foreignKey:SubjectID" binding:"-" json:"-"`
}
