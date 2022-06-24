package tests

import (
	"github.com/stretchr/testify/assert"
	"studentOtomasyon/common"
	"studentOtomasyon/models"
	"testing"
)

func TestCreateLesson(t *testing.T) {
	dbConn := common.GetDbConn()
	lessonID, err := common.CreateLesson(models.Lesson{
		Title:             "Felsefe",
		Quota:             50,
		LessonTime:        "Cuma 12:00",
		TeacherIdOfLesson: 3,
	}, dbConn)
	assert.NoError(t, err, "Err should be nil")
	assert.Greater(t, lessonID, int64(0), "LessonID should be grater than 0")
}

func TestGetLesson(t *testing.T) {
	dbConn := common.GetDbConn()
	respData, err := common.GetLesson(22, dbConn)
	if err != nil {
	}
	assert.Equal(t, "Felsefe", respData.Title, "Lesson name should be Felsefe")
}

func TestDeleteLesson(t *testing.T) {
	dbConn := common.GetDbConn()
	err := common.DeleteLesson(22, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestUpdateLesson(t *testing.T) {
	dbConn := common.GetDbConn()
	err := common.UpdateLesson(models.Lesson{
		ID:                21,
		Title:             "Felsefe",
		Quota:             30,
		LessonTime:        "Cuma 11:00",
		TeacherIdOfLesson: 3,
	}, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestListLesson(t *testing.T) {
	dbConn := common.GetDbConn()
	models, err := common.ListLesson(dbConn)
	length := len(models)
	assert.NoError(t, err)
	assert.Equal(t, 5, length)
}

func TestTitleForLessonID(t *testing.T) {
	dbConn := common.GetDbConn()
	var title string
	title = "Felsefe"
	LessonID, err := common.TitleForLessonID(title, dbConn)
	assert.NoError(t, err, "Err should be nil")
	assert.Greater(t, LessonID, int64(0), "LessonID should greater than 0")
}

func TestLessonIDForGrades(t *testing.T) {
	dbConn := common.GetDbConn()
	var lessonID int64
	lessonID = 5
	models, err := common.LessonIDForGrades(lessonID, dbConn)
	length := len(models)
	assert.Equal(t, 4, length)
	assert.NoError(t, err, "Err should be nil")
}

func TestLessonIDForJustGrades(t *testing.T) {
	dbConn := common.GetDbConn()
	var lessonID int64
	lessonID = 5
	models, err := common.LessonIDForJustGrades(lessonID, dbConn)
	length := len(models)
	assert.Equal(t, 4, length)
	assert.NoError(t, err, "Err should be nil")
}

func TestAddingStudentIDAndLessonIDToLessonStudent(t *testing.T) {
	dbConn := common.GetDbConn()
	var studentID int64
	var lessonID int64
	studentID = 12
	lessonID = 7
	err := common.AddingStudentIDAndLessonIDToLessonStudent(studentID, lessonID, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestDeleteFromLessonStudent(t *testing.T) {
	dbConn := common.GetDbConn()
	var studentID int64
	var lessonID int64
	studentID = 12
	lessonID = 7
	err := common.DeleteFromLessonStudent(studentID, lessonID, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestIsExists(t *testing.T) {
	dbConn := common.GetDbConn()
	var studentID int64
	var lessonID int64
	isExistLesson, err := common.IsExists(studentID, lessonID, dbConn)
	assert.NoError(t, err, "Err should be nil")
	assert.Equal(t, false, isExistLesson)
}

func TestInsertNotes(t *testing.T) {
	dbConn := common.GetDbConn()
	var studentID int64
	var lessonID int64
	var TypeOfNote string
	var ValueOfNote int64
	studentID = 12
	lessonID = 9
	TypeOfNote = "First Exam"
	ValueOfNote = 100
	err := common.InsertNotes(studentID, lessonID, TypeOfNote, ValueOfNote, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestStudentIDAndLessonIDForGrades(t *testing.T) {
	dbConn := common.GetDbConn()
	var studentID int64
	var lessonID int64
	studentID = 12
	lessonID = 9
	respData, err := common.StudentIDAndLessonIDForGrades(studentID, lessonID, dbConn)
	if err != nil {
	}
	ortalama := (respData.FirstExam + respData.SecondExam + respData.FirstPresentation + respData.SecondPresentation) / 4
	assert.Greater(t, int64(93), ortalama)
}
