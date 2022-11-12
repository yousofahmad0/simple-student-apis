package repository

import (
	"gorm.io/gorm"
	"task_2/entity"
)

type StudentRepository struct {
	DB *gorm.DB
}

// GetAll reminders
func (s StudentRepository) GetAll() []entity.Student {
	var students []entity.Student
	_ = s.DB.Find(&students)
	return students
}
