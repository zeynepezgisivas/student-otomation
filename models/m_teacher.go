package models

import (
	"database/sql"
	"log"
)

type Teacher struct {
	ID        int64  `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	UserName  string `db:"username"`
	Password  string `db:"password"`
	Degree    string `db:"degree"`
}

func (t Teacher) Create(query string, dbConn *sql.DB) (int64, error) {
	var newID int64

	hashedPassword, err := HashPassword(t.Password)
	if err != nil {
		return 0, err
	}

	sq := `insert into teacher (first_name, last_name, username, password, degree) values ($1, $2, $3, $4, $5) returning id`
	err = dbConn.QueryRow(sq, t.FirstName, t.LastName, t.UserName, hashedPassword, t.Degree).Scan(&newID)
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	return newID, nil
}

func (t Teacher) Get(ID int64, query string, dbConn *sql.DB) error {

	sq := `select first_name, last_name, username, password, degree from teacher where id = $1`
	err := dbConn.QueryRow(sq, t.ID).Scan(&t.FirstName, &t.LastName, &t.UserName, &t.Password, &t.Degree)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (t Teacher) Delete(ID int64, query string, dbConn *sql.DB) error {

	sq := `delete from teacher where id = $1`
	err := dbConn.QueryRow(sq, t.ID).Err()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (t Teacher) Update(ID int64, query string, dbConn *sql.DB) error {

	sq := `update teacher set first_name=$1, last_name=$2, username=$3, password=$4, degree=$5 where id = $6`
	err := dbConn.QueryRow(sq, t.FirstName, t.LastName, t.UserName, t.Password, t.Degree, t.ID).Err()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (t Teacher) List(query string, dbConn *sql.DB) (interface{}, error) {
	var teacher Teacher
	var teachers []Teacher

	sq := `select id, first_name, last_name, username, password, degree from teacher`
	rows, err := dbConn.Query(sq)
	if err != nil {
		return teachers, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.UserName, &teacher.Password, &teacher.Degree)
		if err != nil {
			log.Fatal(err)
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return teachers, err
}
