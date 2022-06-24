package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"studentOtomasyon/models"
)

func GetDbConn() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=postgres password=tayitkan dbname=student_otomasyon"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func BodyToJson(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return nil
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateStudent(student models.Student, dbConn *sql.DB) (int64, error) {
	var newID int64

	hashedPassword, err := HashPassword(student.Password)
	if err != nil {
		return 0, err
	}

	sq := `insert into student (first_name, last_name, school_number, username, password) values ($1, $2, $3, $4, $5) returning id`
	err = dbConn.QueryRow(sq, student.FirstName, student.LastName, student.SchoolNumber, student.Username, hashedPassword).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func GetStudent(id int64, dbConn *sql.DB) (models.Student, error) {
	var student models.Student

	sq := `select first_name, last_name, school_number, username from student where id = $1`
	err := dbConn.QueryRow(sq, id).Scan(&student.FirstName, &student.LastName, &student.SchoolNumber, &student.Username)
	if err != nil {
		return student, err
	}
	defer dbConn.Close()

	return student, nil
}

func DeleteStudent(id int64, dbConn *sql.DB) error {

	sq := `delete from student where id = $1`
	err := dbConn.QueryRow(sq, id).Err()
	if err != nil {
		return err
	}

	return nil
}

func UpdateStudent(student models.Student, dbConn *sql.DB) error {

	sq := `update student set first_name=$1, last_name=$2, school_number=$3, username=$4, password=$5 where id = $6`
	_, err := dbConn.Exec(sq, student.FirstName, student.LastName, student.SchoolNumber, student.Username, student.Password, student.ID)
	if err != nil {
		return err
	}

	return nil
}

func ListStudent(dbConn *sql.DB) ([]models.Student, error) {
	var student models.Student
	var students []models.Student

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
}

func SchoolNumberForStudentID(school_number int64, dbConn *sql.DB) (int64, error) {
	var studentID int64

	sq := `select id from student where school_number = $1`
	err := dbConn.QueryRow(sq, school_number).Scan(&studentID)
	if err != nil {
		return 0, err
	}

	return studentID, nil
}

func StudentIDForGrades(studentID int64, dbConn *sql.DB) ([]models.Grades, error) {
	var grades []models.Grades
	var grade models.Grades

	sq := `select first_exam
     			, second_exam
     			, first_presentation
     			, second_presentation
     			, lesson_id
     			, (select title from lesson where lesson.id = grades.lesson_id) as lesson_title 
				, student_id
				, (select username from student where student.id = grades.student_id) as student_username
       from grades where student_id = $1`
	rows, err := dbConn.Query(sq, studentID)
	if err != nil {
		return grades, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&grade.FirstExam, &grade.SecondExam, &grade.FirstPresentation, &grade.SecondPresentation, &grade.LessonId, &grade.LessonTitle, &grade.StudentId, &grade.StudentUserName)
		if err != nil {
			log.Fatal(err)
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func CreateTeacher(teacher models.Teacher, dbConn *sql.DB) (int64, error) {
	var newID int64

	hashedPassword, err := HashPassword(teacher.Password)
	if err != nil {
		return 0, err
	}

	sq := `insert into teacher (first_name, last_name, username, password, degree) values ($1, $2, $3, $4, $5) returning id`
	err = dbConn.QueryRow(sq, teacher.FirstName, teacher.LastName, teacher.UserName, hashedPassword, teacher.Degree).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func GetTeacher(id int64, dbConn *sql.DB) (models.Teacher, error) {
	var teacher models.Teacher

	sq := `select first_name, last_name, username, password, degree from teacher where id = $1`
	err := dbConn.QueryRow(sq, id).Scan(&teacher.FirstName, &teacher.LastName, &teacher.UserName, &teacher.Password, &teacher.Degree)
	if err != nil {
		return teacher, err
	}

	return teacher, nil
}

func DeleteTeacher(id int64, dbConn *sql.DB) error {

	sq := `delete from teacher where id = $1`
	err := dbConn.QueryRow(sq, id).Err()
	if err != nil {
		return err
	}

	return nil
}

func UpdateTeacher(teacher models.Teacher, dbConn *sql.DB) error {

	sq := `update teacher set first_name=$1, last_name=$2, username=$3, password=$4, degree=$5 where id = $6`
	err := dbConn.QueryRow(sq, teacher.FirstName, teacher.LastName, teacher.UserName, teacher.Password, teacher.Degree, teacher.ID).Err()
	if err != nil {
		return err
	}

	return nil
}

func ListTeacher(dbConn *sql.DB) ([]models.Teacher, error) {
	var teacher models.Teacher
	var teachers []models.Teacher

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
}

func UserNameForTeacherID(username string, dbConn *sql.DB) (int64, error) {
	var teacherID int64

	sq := `select id from teacher where username = $1`
	err := dbConn.QueryRow(sq, username).Scan(&teacherID)
	if err != nil {
		return 0, err
	}

	return teacherID, nil
}

func TeacherIDOfLessonForLessonID(teacher_id_of_lesson int64, dbConn *sql.DB) ([]int64, error) {
	var lessonIDs []int64
	var lessonID int64

	sq := `select id from lesson where teacher_id_of_lesson = $1`
	rows, err := dbConn.Query(sq, teacher_id_of_lesson)
	if err != nil {
		return lessonIDs, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&lessonID)
		if err != nil {
			log.Fatal(err)
		}
		lessonIDs = append(lessonIDs, lessonID)
	}
	return lessonIDs, nil
}

func TeachersGrades(lessonIDs []int64, dbConn *sql.DB) ([]models.Grades, error) {
	var grades []models.Grades
	var grade models.Grades

	whereClause := "where lesson_id in"
	for i, row := range lessonIDs {
		if i == 0 {
			whereClause = fmt.Sprintf("%s (%d", whereClause, row)
		} else {
			whereClause = fmt.Sprintf("%s, %d", whereClause, row)
		}
	}
	whereClause = fmt.Sprintf("%s)", whereClause)

	sq := fmt.Sprintf(`select first_exam, second_exam, first_presentation, second_presentation from grades %s`, whereClause)
	rowsGrades, err := dbConn.Query(sq)
	if err != nil {
		log.Fatal(err)
		return grades, err
	}

	for rowsGrades.Next() {
		err := rowsGrades.Scan(&grade.FirstExam, &grade.SecondExam, &grade.FirstPresentation, &grade.SecondPresentation)
		grades = append(grades, grade)
		if err != nil {
			log.Fatal(err)
			return grades, nil
		}
	}

	return grades, nil
}

func LessonsAverage(grades []models.Grades) int64 {
	var TotalOfFirstExam int64
	var TotalOfSecondExam int64
	var TotalOfFirstPresentation int64
	var TotalOfSecondPresentation int64
	var lessonsaverage int64

	length := len(grades)

	for _, row := range grades {
		TotalOfFirstExam = TotalOfFirstExam + row.FirstExam
		TotalOfSecondExam = TotalOfSecondExam + row.SecondExam
		TotalOfFirstPresentation = TotalOfFirstPresentation + row.FirstPresentation
		TotalOfSecondPresentation = TotalOfSecondPresentation + row.SecondPresentation
	}

	TotalOfFirstExam = TotalOfFirstExam / int64(length)
	TotalOfSecondExam = TotalOfSecondExam / int64(length)
	TotalOfFirstPresentation = TotalOfFirstPresentation / int64(length)
	TotalOfSecondPresentation = TotalOfSecondPresentation / int64(length)

	lessonsaverage = (TotalOfFirstExam + TotalOfSecondExam + TotalOfFirstPresentation + TotalOfSecondPresentation) / 4

	return lessonsaverage
}

func TeachersLessons(lessonIDs []int64, dbConn *sql.DB) ([]models.Grades, error) {
	var grades []models.Grades
	var grade models.Grades

	whereClause := "where lesson_id in"
	for i, row := range lessonIDs {
		if i == 0 {
			whereClause = fmt.Sprintf("%s (%d", whereClause, row)
		} else {
			whereClause = fmt.Sprintf("%s, %d", whereClause, row)
		}
	}
	whereClause = fmt.Sprintf("%s)", whereClause)

	sq := fmt.Sprintf(`select first_exam, second_exam, first_presentation, second_presentation from grades %s`, whereClause)
	rows, err := dbConn.Query(sq)
	if err != nil {
		log.Fatal(err)
		return grades, err
	}

	for rows.Next() {
		err := rows.Scan(&grade.FirstExam, &grade.SecondExam, &grade.FirstPresentation, &grade.SecondPresentation)
		grades = append(grades, grade)
		if err != nil {
			log.Fatal(err)
			return grades, err
		}
	}
	return grades, nil
}

func TeacherIDOfLessonForLessonInformation(teacher_id_of_lesson int64, dbConn *sql.DB) ([]models.Lesson, error) {
	var lessons []models.Lesson
	var lesson models.Lesson

	sq := `select id, title, quota, lesson_time from lesson where teacher_id_of_lesson = $1`
	row, err := dbConn.Query(sq, teacher_id_of_lesson)
	if err != nil {
		log.Fatal(err)
		return lessons, err
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&lesson.ID, &lesson.Title, &lesson.Quota, &lesson.LessonTime)
		lessons = append(lessons, lesson)
		if err != nil {
			log.Fatal(err)
			return lessons, err
		}
	}
	return lessons, nil
}

func CreateLesson(lesson models.Lesson, dbConn *sql.DB) (int64, error) {
	var newID int64

	sq := `insert into lesson (title, quota, lesson_time, teacher_id_of_lesson) values ($1, $2, $3, $4) returning id`
	err := dbConn.QueryRow(sq, lesson.Title, lesson.Quota, lesson.LessonTime, lesson.TeacherIdOfLesson).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

func GetLesson(id int64, dbConn *sql.DB) (models.Lesson, error) {
	var lesson models.Lesson

	sq := `select title, quota, lesson_time, teacher_id_of_lesson from lesson where id = $1`
	err := dbConn.QueryRow(sq, id).Scan(&lesson.Title, &lesson.Quota, &lesson.LessonTime, &lesson.TeacherIdOfLesson)
	if err != nil {
		return lesson, err
	}

	return lesson, nil
}

func DeleteLesson(id int64, dbConn *sql.DB) error {

	sq := `delete from lesson where id = $1`
	err := dbConn.QueryRow(sq, id).Err()
	if err != nil {
		return err
	}

	return nil
}

func UpdateLesson(lesson models.Lesson, dbConn *sql.DB) error {

	sq := `update lesson set title=$1, quota=$2, lesson_time=$3, teacher_id_of_lesson=$4 where id=$5`
	err := dbConn.QueryRow(sq, lesson.Title, lesson.Quota, lesson.LessonTime, lesson.TeacherIdOfLesson, lesson.ID).Err()
	if err != nil {
		return err
	}

	return nil
}

func ListLesson(dbConn *sql.DB) ([]models.Lesson, error) {
	var lesson models.Lesson
	var lessons []models.Lesson

	sq := `select title, quota, lesson_time, teacher_id_of_lesson from lesson`
	rows, err := dbConn.Query(sq)
	if err != nil {
		return lessons, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&lesson.Title, &lesson.Quota, &lesson.LessonTime, &lesson.TeacherIdOfLesson)
		if err != nil {
			log.Fatal(err)
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

func TitleForLessonID(title string, dbConn *sql.DB) (int64, error) {
	var lessonID int64

	sq := `select id from lesson where title = $1`
	err := dbConn.QueryRow(sq, title).Scan(&lessonID)
	if err != nil {
		return 0, err
	}
	return lessonID, nil
}

func LessonIDForGrades(lessonID int64, dbConn *sql.DB) ([]models.Grades, error) {
	var grade models.Grades
	var grades []models.Grades

	sq := `select first_exam, second_exam, first_presentation, second_presentation, student_id, lesson_id from grades where lesson_id = $1`
	row, err := dbConn.Query(sq, lessonID)
	if err != nil {
		log.Fatal(err)
		return grades, err
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&grade.FirstExam, &grade.SecondExam, &grade.FirstPresentation, &grade.SecondPresentation, &grade.StudentId, &grade.LessonId)
		grades = append(grades, grade)
		if err != nil {
			log.Fatal(err)
		}
	}
	return grades, nil
}

func LessonIDForJustGrades(lessonID int64, dbConn *sql.DB) ([]models.Grades, error) {
	var grades []models.Grades

	sq := `select first_exam, second_exam, first_presentation, second_presentation from grades where lesson_id = $1`
	row, err := dbConn.Query(sq, lessonID)
	if err != nil {
		log.Fatal(err)
		return grades, err
	}
	defer row.Close()

	for row.Next() {
		var grade models.Grades
		err := row.Scan(&grade.FirstExam, &grade.SecondExam, &grade.FirstPresentation, &grade.SecondPresentation)
		if err != nil {
			log.Fatal(err)
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func LessonsAverages(grades []models.Grades) int64 {
	var totalOfFirstExam int64
	var totalOfSecondExam int64
	var totalOfFirstPresentation int64
	var totalOfSecondPresentation int64
	var average int64

	length := len(grades)

	for _, row := range grades {
		totalOfFirstExam = totalOfFirstExam + row.FirstExam
		totalOfSecondExam = totalOfSecondExam + row.SecondExam
		totalOfFirstPresentation = totalOfFirstPresentation + row.FirstPresentation
		totalOfSecondPresentation = totalOfSecondPresentation + row.SecondPresentation
	}

	totalOfFirstExam = totalOfFirstExam / int64(length)
	totalOfSecondExam = totalOfSecondExam / int64(length)
	totalOfFirstPresentation = totalOfFirstPresentation / int64(length)
	totalOfSecondPresentation = totalOfSecondPresentation / int64(length)

	average = (totalOfFirstExam + totalOfSecondExam + totalOfFirstPresentation + totalOfSecondPresentation) / 4

	return average
}

func AddingStudentIDAndLessonIDToLessonStudent(studentID int64, lessonID int64, dbConn *sql.DB) error {

	sq := `insert into lesson_student (student_id, lesson_id) values ($1, $2)`
	err := dbConn.QueryRow(sq, studentID, lessonID).Err()
	if err != nil {
		return err
	}

	return nil
}

func DeleteFromLessonStudent(studentID int64, lessonID int64, dbConn *sql.DB) error {

	sq := `delete from lesson_student where student_id = $1 and lesson_id = $2`
	err := dbConn.QueryRow(sq, studentID, lessonID).Err()
	if err != nil {
		return err
	}

	return nil
}

func StudentIDAndLessonIDForGrades(studentID int64, lessonID int64, dbConn *sql.DB) (models.Grades, error) {
	var grade models.Grades

	sq := `select first_exam, second_exam, first_presentation, second_presentation from grades where student_id = $1 and lesson_id = $2`
	err := dbConn.QueryRow(sq, studentID, lessonID).Scan(&grade.FirstExam, &grade.SecondExam, &grade.FirstPresentation, &grade.SecondPresentation)
	if err != nil {
		log.Fatal(err)
	}

	return grade, nil
}

func IsExists(studentID int64, lessonID int64, dbConn *sql.DB) (bool, error) {
	var isexists bool
	var bool bool

	sq := "select exists(select 1 from lesson_student where student_id=$1 and lesson_id=$2)"
	err := dbConn.QueryRow(sq, studentID, lessonID).Scan(&isexists)
	if err != nil {
		log.Fatal(err)
		return isexists, err
	}

	if isexists == false {
		log.Print("Öğrenci bu dersi almıyor.")
		return bool, err
	}

	return bool, err
}

func InsertNotes(StudentID int64, LessonID int64, TypeOfNote string, ValueOfNote int64, dbConn *sql.DB) error {
	var isexists bool

	sq := `select exists(select 1 from grades where student_id=$1 and lesson_id=$2)`
	err := dbConn.QueryRow(sq, StudentID, LessonID).Scan(&isexists)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if isexists == false {
		var insertSq string
		switch TypeOfNote {
		case "First Exam":
			insertSq = `insert into grades (first_exam, student_id, lesson_id) values ($1, $2, $3)`
		case "Second Exam":
			insertSq = `insert into grades (second_exam, student_id, lesson_id) values ($1, $2, $3)`
		case "First Presentation":
			insertSq = `insert into grades (first_presentation, student_id, lesson_id) values ($1, $2, $3)`
		case "Second Presentation":
			insertSq = `insert into grades (second_presentation, student_id, lesson_id) values ($1, $2, $3)`
		}
		err = dbConn.QueryRow(insertSq, ValueOfNote, StudentID, LessonID).Err()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	if isexists == true {
		var updateSq string
		switch TypeOfNote {
		case "First Exam":
			updateSq = `update grades set first_exam = $1 where student_id = $2 and lesson_id = $3`
		case "Second Exam":
			updateSq = `update grades set second_exam = $1 where student_id = $2 and lesson_id = $3`
		case "First Presentation":
			updateSq = `update grades set first_presentation = $1 where student_id = $2 and lesson_id = $3`
		case "Second Presentation":
			updateSq = `update grades set second_presentation = $1 where student_id = $2 and lesson_id = $3`
		}
		err = dbConn.QueryRow(updateSq, ValueOfNote, StudentID, LessonID).Err()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
