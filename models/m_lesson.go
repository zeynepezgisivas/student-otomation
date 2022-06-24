package models

import (
	"database/sql"
	"log"
)

type Lesson struct {
	ID                int64  `db:"id"`
	Title             string `db:"title"`
	Quota             int64  `db:"quota"`
	LessonTime        string `db:"lesson_time"`
	TeacherIdOfLesson int64  `db:"teacher_id_of_lesson"`
}

func (l Lesson) Create(query string, dbConn *sql.DB) (int64, error) {
	var newID int64

	sq := `insert into lesson (title, quota, lesson_time, teacher_id_of_lesson) values ($1, $2, $3, $4) returning id`
	err := dbConn.QueryRow(sq, l.Title, l.Quota, l.LessonTime, l.TeacherIdOfLesson).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func (l Lesson) Get(ID int64, query string, dbConn *sql.DB) error {

	sq := `select title, quota, lesson_time, teacher_id_of_lesson from lesson where id = $1`
	err := dbConn.QueryRow(sq, l.ID).Scan(&l.Title, &l.Quota, &l.LessonTime, &l.TeacherIdOfLesson)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (l Lesson) Delete(ID int64, query string, dbConn *sql.DB) error {
	sq := `delete from lesson where id = $1`
	err := dbConn.QueryRow(sq, l.ID).Err()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (l Lesson) Update(ID int64, query string, dbConn *sql.DB) error {
	sq := `update lesson set title=$1, quota=$2, lesson_time=$3, teacher_id_of_lesson=$4 where id=$5`
	err := dbConn.QueryRow(sq, l.Title, l.Quota, l.LessonTime, l.TeacherIdOfLesson, l.ID).Err()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return err
}

func (l Lesson) List(query string, dbConn *sql.DB) (interface{}, error) {
	var lesson Lesson
	var lessons []Lesson

	sq := `select title, quota, lesson_time, teacher_id_of_lesson from lesson`
	rows, err := dbConn.Query(sq)
	if err != nil {
		return lessons, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&l.Title, &l.Quota, &l.LessonTime, &l.TeacherIdOfLesson)
		if err != nil {
			log.Fatal(err)
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return lessons, err
}

type Grades struct {
	ID                        int64  `db:"id"`
	FirstExam                 int64  `db:"first_exam"`
	SecondExam                int64  `db:"second_exam"`
	FirstPresentation         int64  `db:"first_presentation"`
	SecondPresentation        int64  `db:"second_presentation"`
	StudentId                 int64  `db:"student_id"`
	StudentUserName           string `db:"student_username"`
	LessonId                  int64  `db:"lesson_id"`
	LessonTitle               string `db:"lesson_title"`
	TotalOfFirstExam          int64  `db:"total_of_first_exam"`
	TotalOfSecondExam         int64  `db:"total_of_second_exam"`
	TotalOfFirstPresentation  int64  `db:"total_of_first_presentation"`
	TotalOfSecondPresentation int64  `db:"total_of_second_presentation"`
}

type LessonStudent struct {
	ID        int64 `db:"id"`
	StudentId int64 `db:"student_id"`
	LessonId  int64 `db:"lesson_id"`
}
