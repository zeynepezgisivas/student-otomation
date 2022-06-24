package tests

import (
	"github.com/stretchr/testify/assert"
	"studentOtomasyon/common"
	"studentOtomasyon/models"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	dbConn := common.GetDbConn()
	studentID, err := common.CreateStudent(models.Student{
		FirstName:    "Zeynep",
		LastName:     "Yaprak",
		SchoolNumber: 2022110,
		Username:     "zeynepyaprak",
		Password:     "zeynepyaprak1234",
	}, dbConn)
	assert.NoError(t, err, "Err should be nil")
	assert.Greater(t, studentID, int64(0), "StudentID should greater than 0")
}

func TestGetStudent(t *testing.T) {
	dbConn := common.GetDbConn()
	respData, err := common.GetStudent(17, dbConn)
	if err != nil {
	}
	assert.Equal(t, "Zeynep", respData.FirstName, "Student name should be Zeynep")
}

func TestDeleteStudent(t *testing.T) {
	dbConn := common.GetDbConn()
	err := common.DeleteStudent(17, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestUpdateStudent(t *testing.T) {
	dbConn := common.GetDbConn()
	err := common.UpdateStudent(models.Student{
		ID:           5,
		FirstName:    "Ali",
		LastName:     "Yılmaz",
		SchoolNumber: 2018110,
		Username:     "aliyilmazz",
		Password:     "",
	}, dbConn)
	assert.NoError(t, err, "Err should be nil")
}

func TestListStudent(t *testing.T) {
	dbConn := common.GetDbConn()
	models, err := common.ListStudent(dbConn)
	length := len(models)
	assert.NoError(t, err)
	assert.Equal(t, 5, length)
}

func TestSchoolNumberForStudentID(t *testing.T) {
	dbConn := common.GetDbConn()
	var school_number int64
	school_number = 2020110
	studentID, err := common.SchoolNumberForStudentID(school_number, dbConn)
	assert.NoError(t, err, "Err should be nil")
	assert.Greater(t, studentID, int64(0), "StudentID should greater than 0")
}

func TestStudentIDForGrades(t *testing.T) {
	dbConn := common.GetDbConn()
	var studentID int64
	studentID = 8
	models, err := common.StudentIDForGrades(studentID, dbConn)
	length := len(models)
	assert.Equal(t, 4, length)
	assert.NoError(t, err, "Err should be nil")
}
