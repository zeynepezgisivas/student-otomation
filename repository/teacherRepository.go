package repository

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"studentOtomasyon/models"
)

type TRepository interface {
	CreateTeacher(teacher models.Teacher) error
	GetTeacher(id int64) (*models.Teacher, error)
	DeleteTeacher(id int64) error
	UpdateTeacher(teacher models.Teacher) error
	ListTeacher() ([]models.Teacher, error)
}

type TeacherRepository struct {
	connection *sqlx.DB
}

func NewTeacherRepository(db *sqlx.DB) *TeacherRepository {
	return &TeacherRepository{connection: db}
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (repo *TeacherRepository) CreateTeacher(teacher models.Teacher) error {
	var newID int64

	hashedPassword, err := HashPassword(teacher.Password)
	if err != nil {
		return err
	}

	sq := `insert into teacher (first_name, last_name, username, password, degree) values ($1, $2, $3, $4, $5) returning id`
	repo.connection.QueryRowx(sq, teacher.FirstName, teacher.LastName, teacher.UserName, hashedPassword, teacher.Degree).Scan(&newID)
	return nil
}

func (repo *TeacherRepository) GetTeacher(id int64) (*models.Teacher, error) {
	var teacher models.Teacher

	sq := `select id,first_name, last_name, username, password, degree from teacher where id = $1`
	err := repo.connection.Get(&teacher, sq, id)
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (repo TeacherRepository) DeleteTeacher(id int64) error {

	sq := `delete from teacher where id = $1`
	_, err := repo.connection.Exec(sq, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TeacherRepository) UpdateTeacher(teacher models.Teacher) error {

	sq := `update teacher set first_name=$1, last_name=$2, username=$3, password=$4, degree=$5 where id = $6`
	_, err := repo.connection.Exec(sq, teacher.FirstName, teacher.LastName, teacher.UserName, teacher.Password, teacher.Degree, teacher.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TeacherRepository) ListTeacher() ([]models.Teacher, error) {
	var teachers []models.Teacher

	sq := `select id, first_name, last_name, username, password, degree from teacher`
	err := repo.connection.Select(&teachers, sq)
	if err != nil {
		return nil, err
	}

	return teachers, err
}
