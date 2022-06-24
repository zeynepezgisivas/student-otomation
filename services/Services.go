package services

import (
	"github.com/jmoiron/sqlx"
	"studentOtomasyon/repository"
)

type Services struct {
	LessonService  *LessonService
	StudentService *StudentService
	TeacherService *TeacherService
}

func GetServices(db *sqlx.DB) *Services {
	repositories := repository.GetRepositories(db)
	return &Services{
		LessonService:  &LessonService{Repository: repositories.LessonRepo},
		StudentService: &StudentService{Repository: repositories.StudentRepo},
		TeacherService: &TeacherService{Repository: repositories.TeacherRepo},
	}
}
