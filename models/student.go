package models

type Student struct {
	ID    uint   `gorm:"primaryKey" json:"student_id"`
	Name  string `json:"name" binding:"required"`
	Group string `json:"group" binding:"required"`
	Email string `gorm:"unique" json:"email" binding:"required,email"`
}
