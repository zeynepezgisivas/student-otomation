package services

import (
	"log"
	"studentOtomasyon/models"
	"studentOtomasyon/repository"
)

type SService interface {
	CreateStudent(student models.Student) error
	GetStudent(id int64) (*models.Student, error)
	DeleteStudent(id int64) error
	UpdateStudent(student models.Student) error
	ListStudent() ([]models.Student, error)
}

type StudentService struct {
	Repository *repository.StudentRepository
}

func (ser *StudentService) CreateStudent(student models.Student) error {

	err := ser.Repository.CreateStudent(student)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *StudentService) GetStudent(id int64) (*models.Student, error) {
	var student *models.Student

	student, err := ser.Repository.GetStudent(id)
	if err != nil {
		log.Fatal(err)
		return student, err
	}
	return student, nil
}

func (ser *StudentService) DeleteStudent(id int64) error {

	err := ser.Repository.DeleteStudent(id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *StudentService) UpdateStudent(student models.Student) error {

	err := ser.Repository.UpdateStudent(student)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *StudentService) ListStudent() ([]models.Student, error) {
	var students []models.Student

	students, err := ser.Repository.ListStudent()
	if err != nil {
		log.Fatal(err)
		return students, err
	}
	return students, err
}
