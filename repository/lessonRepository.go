package repository

import (
	"github.com/jmoiron/sqlx"
	"studentOtomasyon/models"
)

type Repository interface {
	CreateLesson(lesson models.Lesson) models.Lesson
	GetLesson(id int64) (*models.Lesson, error)
	DeleteLesson(id int64) error
	UpdateLesson(lesson models.Lesson) error
	ListLesson() ([]models.Lesson, error)
}

type LessonRepository struct {
	connection *sqlx.DB
}

func NewLessonRepository(db *sqlx.DB) *LessonRepository {
	return &LessonRepository{
		connection: db,
	}
}

func (repo *LessonRepository) CreateLesson(lesson models.Lesson) error {
	var newID int64

	sq := `insert into lesson (title, quota, lesson_time, teacher_id_of_lesson) values ($1, $2, $3, $4) returning id`
	repo.connection.QueryRow(sq, lesson.Title, lesson.Quota, lesson.LessonTime, lesson.TeacherIdOfLesson).Scan(&newID)
	return nil
}

func (repo *LessonRepository) GetLesson(id int64) (*models.Lesson, error) {
	var model models.Lesson

	sq := `select title, quota, lesson_time, teacher_id_of_lesson from lesson where id = $1`
	err := repo.connection.Get(&model, sq, id)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (repo *LessonRepository) DeleteLesson(id int64) error {

	sq := `delete from lesson where id = $1`
	_, err := repo.connection.Exec(sq, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *LessonRepository) UpdateLesson(lesson models.Lesson) error {

	sq := `update lesson set title=$1, quota=$2, lesson_time=$3, teacher_id_of_lesson=$4 where id=$5`
	_, err := repo.connection.Exec(sq, lesson.Title, lesson.Quota, lesson.LessonTime, lesson.TeacherIdOfLesson, lesson.ID)
	if err != nil {
		return nil
	}

	return err
}

func (repo *LessonRepository) ListLesson() ([]models.Lesson, error) {
	var lessons []models.Lesson

	sq := `select id, title, quota, lesson_time, teacher_id_of_lesson from lesson`
	err := repo.connection.Select(&lessons, sq)
	if err != nil {
		return nil, err
	}

	return lessons, err
}
