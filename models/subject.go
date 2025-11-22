package models

type Subject struct {
	ID   uint   `gorm:"primaryKey" json:"subject_id"`
	Name string `json:"name" binding:"required"`
}
