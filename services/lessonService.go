package services

import (
	"log"
	"studentOtomasyon/models"
	"studentOtomasyon/repository"
)

type Service interface {
	CreateLesson(lesson models.Lesson) models.Lesson
	GetLesson(id int64) (*models.Lesson, error)
	DeleteLesson(id int64) error
	UpdateLesson(lesson models.Lesson) error
	ListLesson() ([]models.Lesson, error)
}

type LessonService struct {
	Repository *repository.LessonRepository
}

func (ser *LessonService) CreateLesson(lesson models.Lesson) error {

	err := ser.Repository.CreateLesson(lesson)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *LessonService) GetLesson(id int64) (*models.Lesson, error) {
	var lesson *models.Lesson

	lesson, err := ser.Repository.GetLesson(id)
	if err != nil {
		log.Fatal(err)
		return lesson, err
	}
	return lesson, nil
}

func (ser *LessonService) DeleteLesson(id int64) error {

	err := ser.Repository.DeleteLesson(id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *LessonService) UpdateLesson(lesson models.Lesson) error {

	err := ser.Repository.UpdateLesson(lesson)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (ser *LessonService) ListLesson() ([]models.Lesson, error) {
	var lessons []models.Lesson

	lessons, err := ser.Repository.ListLesson()
	if err != nil {
		log.Fatal(err)
		return lessons, err
	}
	return lessons, err
}
