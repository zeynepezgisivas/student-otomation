package api

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"studentOtomasyon/common"
	mhttp "studentOtomasyon/http"
	"studentOtomasyon/models"
	"studentOtomasyon/services"
	"studentOtomasyon/svc"
)

type StudentApi struct {
	Service *services.StudentService
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Student

	err := common.BodyToJsonReq(r, &reqBody)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = services.GetServices(svc.App.DB).StudentService.CreateStudent(reqBody)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	data, err := services.GetServices(svc.App.DB).StudentService.GetStudent(id)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, data)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	err := services.GetServices(svc.App.DB).StudentService.DeleteStudent(id)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	var reqBody struct {
		FirstName    string
		LastName     string
		SchoolNumber int64
		Username     string
		Password     string
	}
	err := common.BodyToJsonReq(r, &reqBody)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = services.GetServices(svc.App.DB).StudentService.UpdateStudent(models.Student{
		ID:           id,
		FirstName:    reqBody.FirstName,
		LastName:     reqBody.LastName,
		SchoolNumber: reqBody.SchoolNumber,
		Username:     reqBody.Username,
		Password:     reqBody.Password,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func ListStudent(w http.ResponseWriter, r *http.Request) {

	data, err := services.GetServices(svc.App.DB).StudentService.ListStudent()
	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, data)
}
