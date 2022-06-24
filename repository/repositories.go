package repository

import (
	"github.com/jmoiron/sqlx"
	//"github.com/jmoiron/sqlx"
)

type Repositories struct {
	LessonRepo  *LessonRepository
	StudentRepo *StudentRepository
	TeacherRepo *TeacherRepository
}

func GetRepositories(db *sqlx.DB) *Repositories {
	lessonRepo := NewLessonRepository(db)
	studentRepo := NewStudentRepository(db)
	teacherRepo := NewTeacherRepository(db)
	return &Repositories{
		LessonRepo:  lessonRepo,
		StudentRepo: studentRepo,
		TeacherRepo: teacherRepo,
	}
}
