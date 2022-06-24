package tests

import (
	"github.com/stretchr/testify/assert"
	"studentOtomasyon/common"
	"studentOtomasyon/models"
	"testing"
)

func TestCreateTeacher(t *testing.T) {
	dbConn := common.GetDbConn()
	teacherID, err := common.CreateTeacher(models.Teacher{
		FirstName: "Aleyna",
		LastName:  "Sedef",
		UserName:  "aleynasedef",
		Password:  "aleynasedef123456",
		Degree:    "Prof",
	}, dbConn)
	assert.NoError(t, err, "Err should be nil")
	assert.Greater(t, teacherID, int64(0), "TeacherID should greater than 0")
}

func TestGetTeacher(t *testing.T) {
	dbConn := common.GetDbConn()
	respData, err := common.GetTeacher(12, dbConn)
	if err != nil {
	}
	assert.Equal(t, "Aleyna", respData.FirstName, "Teacher name should be aleyna")
}

func TestDeleteTeacher(t *testing.T) {
	dbConn := common.GetDbConn()
	err := common.DeleteTeacher(12, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestUpdateTeacher(t *testing.T) {
	dbConn := common.GetDbConn()
	err := common.UpdateTeacher(models.Teacher{
		ID:        3,
		FirstName: "İlber",
		LastName:  "Ortaylı",
		UserName:  "ilber.ortayli",
		Password:  "",
		Degree:    "Prof",
	}, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestListTeacher(t *testing.T) {
	dbConn := common.GetDbConn()
	models, err := common.ListTeacher(dbConn)
	length := len(models)
	assert.NoError(t, err)
	assert.Equal(t, 4, length)
}

func TestUserNameForTeacherID(t *testing.T) {
	dbConn := common.GetDbConn()
	var username string
	username = "ilber.ortayli"
	TeacherID, err := common.UserNameForTeacherID(username, dbConn)
	assert.NoError(t, err, "Err should be nil")
	assert.Greater(t, TeacherID, int64(0), "TeacherID should greater than 0")
}

func TestTeacherIDOfLessonForLessonID(t *testing.T) {
	dbConn := common.GetDbConn()
	var teacher_id_of_lesson int64
	teacher_id_of_lesson = 3
	models, err := common.TeacherIDOfLessonForLessonID(teacher_id_of_lesson, dbConn)
	length := len(models)
	assert.Equal(t, 2, length)
	assert.NoError(t, err, "Err should be nil")
}

func TestTeachersLessons(t *testing.T) {
	dbConn := common.GetDbConn()
	var lessonIDs []int64
	lessonIDs = []int64{3, 5, 6, 7}
	models, err := common.TeachersLessons(lessonIDs, dbConn)
	length := len(models)
	assert.Equal(t, 8, length)
	assert.NoError(t, err, "Err should be nil")
}

func TestLessonsAverage(t *testing.T) {
	var ortalama int64
	var m []models.Grades
	n := models.Grades{
		FirstExam:          70,
		SecondExam:         80,
		FirstPresentation:  95,
		SecondPresentation: 100,
	}
	z := models.Grades{
		TotalOfFirstExam:          70,
		TotalOfSecondExam:         80,
		TotalOfFirstPresentation:  95,
		TotalOfSecondPresentation: 100,
	}
	w := models.Grades{
		StudentId:   12,
		LessonId:    5,
		LessonTitle: "Matematik",
	}
	m = append(m, n)
	m = append(m, z)
	m = append(m, w)
	ortalama = common.LessonsAverage(m)
	assert.Equal(t, int64(28), ortalama)
}

func TestTeacherIDOfLessonForLessonInformation(t *testing.T) {
	dbConn := common.GetDbConn()
	var teacher_id_of_lesson int64
	models, err := common.TeacherIDOfLessonForLessonInformation(teacher_id_of_lesson, dbConn)
	length := len(models)
	assert.Equal(t, int(0), length)
	assert.NoError(t, err, "Err should be nil")
}

func TestTeachersGrades(t *testing.T) {
	dbConn := common.GetDbConn()
	var lessonIDs []int64
	lessonIDs = []int64{3, 5, 6, 7}
	models, err := common.TeachersGrades(lessonIDs, dbConn)
	length := len(models)
	assert.Equal(t, int(8), length)
	assert.NoError(t, err, "Err should be nil")
}
