package repository

import (
	"fmt"
	"gorm.io/gorm"
	"task_2/entity"
)

type StudentRepository struct {
	DB *gorm.DB
}

// GetAll students
func (s StudentRepository) GetAll() []*entity.Student {
	var students []*entity.Student
	//_ = s.DB.Find(&students)
	if err := s.DB.Find(&students).Error; err != nil {
		return nil
	}
	return students
}

// GetOne student
func (s StudentRepository) GetStudent(id int) (string, *entity.Student) {
	student := new(entity.Student)
	if err := s.DB.First(&student, id).Error; err != nil {
		return "student not found", nil
	}
	return "founded", student
}

// Create student
func (s StudentRepository) Create(u *entity.Student) string {

	result := s.DB.Create(&u)
	fmt.Println(result.Row())
	s.DB.Save(&u)
	return "Student created Successfully"
}

// Delete student
func (s StudentRepository) Delete(id int) string {

	deleted := s.DB.Delete(&entity.Student{}, id)
	if deleted.RowsAffected < 1 {
		return "student not found"
	}

	return "deleted successfully"
}

// Update student
func (s StudentRepository) Update(id int, u *entity.Student) (string, *entity.Student) {

	student := new(entity.Student)
	if err := s.DB.First(&student, id).Error; err != nil {
		return "student not found", nil
	}
	if u.FirstName == "" || u.LastName == "" {
		return "firstname and lastname are required", nil
	}

	student.FirstName = u.FirstName
	student.LastName = u.LastName
	s.DB.Save(student)
	return "student updated successfully", student
}

// Patch Update student
func (s StudentRepository) Patch(id int, u *entity.Student) (string, *entity.Student) {

	student := new(entity.Student)
	if err := s.DB.First(&student, id).Error; err != nil {
		return "student not found", nil
	}

	if u.FirstName != "" {
		student.FirstName = u.FirstName
	}
	if u.LastName != "" {
		student.LastName = u.LastName
	}
	s.DB.Save(student)
	return "student updated", student
}
