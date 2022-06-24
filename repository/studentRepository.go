package repository

import (
	"github.com/jmoiron/sqlx"
	"studentOtomasyon/models"
)

type SRepository interface {
	CreateStudent(student models.Student) error
	GetStudent(id int64) (*models.Student, error)
	DeleteStudent(id int64) error
	UpdateStudent(student models.Student) error
	ListStudent() ([]models.Student, error)
}

type StudentRepository struct {
	connection *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{connection: db}
}

func (repo *StudentRepository) CreateStudent(student models.Student) error {
	var newID int64

	hashedPassword, err := HashPassword(student.Password)
	if err != nil {
		return err
	}

	sq := `insert into student (first_name, last_name, school_number, username, password) values ($1, $2, $3, $4, $5) returning id`
	repo.connection.QueryRowx(sq, student.FirstName, student.LastName, student.SchoolNumber, student.Username, hashedPassword).Scan(&newID)
	return nil
}

func (repo *StudentRepository) GetStudent(id int64) (*models.Student, error) {
	var student models.Student

	sq := `select first_name, last_name, school_number, username from student where id = $1`
	err := repo.connection.Get(&student, sq, id)
	if err != nil {
		return nil, err
	}

	return &student, err
}

func (repo *StudentRepository) DeleteStudent(id int64) error {

	sq := `delete from student where id = $1`
	_, err := repo.connection.Exec(sq, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *StudentRepository) UpdateStudent(student models.Student) error {

	sq := `update student set first_name=$1, last_name=$2, school_number=$3, username=$4, password=$5 where id = $6`
	_, err := repo.connection.Exec(sq, student.FirstName, student.LastName, student.SchoolNumber, student.Username, student.Password, student.ID)
	if err != nil {
		return nil
	}

	return err
}

func (repo *StudentRepository) ListStudent() ([]models.Student, error) {
	var students []models.Student

	sq := `select id, first_name, last_name, school_number, username, password from student`
	err := repo.connection.Select(&students, sq)
	if err != nil {
		return nil, err
	}

	return students, err
}
