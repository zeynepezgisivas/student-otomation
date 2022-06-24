package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Student struct {
	ID           int64  `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	SchoolNumber int64  `db:"school_number"`
	Username     string `db:"username"`
	Password     string `db:"password"`
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s Student) Create(query string, dbConn *sql.DB) (int64, error) {
	var newID int64

	hashedPassword, err := HashPassword(s.Password)
	if err != nil {
		return 0, err
	}

	sq := `insert into student (first_name, last_name, school_number, username, password) values ($1, $2, $3, $4, $5) returning id`
	err = dbConn.QueryRow(sq, s.FirstName, s.LastName, s.SchoolNumber, s.Username, hashedPassword).Scan(&newID)
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	return newID, nil
}

func (s Student) Get(ID int64, query string, dbConn *sql.DB) error {

	sq := `select first_name, last_name, school_number, username from student where id = $1`
	err := dbConn.QueryRow(sq, s.ID).Scan(&s.FirstName, &s.LastName, &s.SchoolNumber, &s.Username)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (s Student) Delete(ID int64, query string, dbConn *sql.DB) error {

	sq := `delete from student where id = $1`
	err := dbConn.QueryRow(sq, s.ID).Err()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (s Student) Update(ID int64, query string, dbConn *sql.DB) error {

	sq := `update student set first_name=$1, last_name=$2, school_number=$3, username=$4, password=$5 where id = $6`
	_, err := dbConn.Exec(sq, s.FirstName, s.LastName, s.SchoolNumber, s.Username, s.Password, s.ID)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (s Student) List(query string, dbConn *sql.DB) (interface{}, error) {
	var student Student
	var students []Student

	sq := `select first_name, last_name, school_number, username, password from student`
	rows, err := dbConn.Query(sq)
	if err != nil {
		return students, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&student.FirstName, &student.LastName, &student.SchoolNumber, &student.Username, &student.Password)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}
	return students, nil
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return students, err
}
