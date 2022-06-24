package services

import (
	"log"
	"studentOtomasyon/models"
	"studentOtomasyon/repository"
)

type TService interface {
	CreateTeacher(teacher models.Teacher) error
	GetTeacher(id int64) (*models.Teacher, error)
	DeleteTeacher(id int64) error
	UpdateTeacher(teacher models.Teacher) error
	ListTeacher() ([]models.Teacher, error)
}

type TeacherService struct {
	Repository *repository.TeacherRepository
}

func (ser *TeacherService) CreateTeacher(teacher models.Teacher) error {

	err := ser.Repository.CreateTeacher(teacher)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *TeacherService) GetTeacher(id int64) (*models.Teacher, error) {
	var teacher *models.Teacher

	teacher, err := ser.Repository.GetTeacher(id)
	if err != nil {
		log.Fatal(err)
		return teacher, err
	}
	return teacher, nil
}

func (ser *TeacherService) DeleteTeacher(id int64) error {

	err := ser.Repository.DeleteTeacher(id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *TeacherService) UpdateTeacher(teacher models.Teacher) error {

	err := ser.Repository.UpdateTeacher(teacher)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *TeacherService) ListTeacher() ([]models.Teacher, error) {
	var teachers []models.Teacher

	teachers, err := ser.Repository.ListTeacher()
	if err != nil {
		log.Fatal(err)
		return teachers, err
	}
	return teachers, err
}
