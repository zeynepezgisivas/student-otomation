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

type TeacherApi struct {
	Service *services.TeacherService
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Teacher

	err := common.BodyToJsonReq(r, &reqBody)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = services.GetServices(svc.App.DB).TeacherService.CreateTeacher(reqBody)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")

}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	data, err := services.GetServices(svc.App.DB).TeacherService.GetTeacher(id)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, data)
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	err := services.GetServices(svc.App.DB).TeacherService.DeleteTeacher(id)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	var reqBody struct {
		FirstName string
		LastName  string
		UserName  string
		Password  string
		Degree    string
	}
	err := common.BodyToJsonReq(r, &reqBody)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = services.GetServices(svc.App.DB).TeacherService.UpdateTeacher(models.Teacher{
		ID:        id,
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		UserName:  reqBody.UserName,
		Password:  reqBody.Password,
		Degree:    reqBody.Degree,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func ListTeacher(w http.ResponseWriter, r *http.Request) {

	data, err := services.GetServices(svc.App.DB).TeacherService.ListTeacher()
	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, data)
}
