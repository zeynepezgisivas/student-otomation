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

type LessonApi struct {
	Service *services.LessonService
}

func CreateLesson(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Lesson

	err := common.BodyToJsonReq(r, &reqBody)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = services.GetServices(svc.App.DB).LessonService.CreateLesson(reqBody)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func GetLesson(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	data, err := services.GetServices(svc.App.DB).LessonService.GetLesson(id)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, data)
}

func DeleteLesson(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	err := services.GetServices(svc.App.DB).LessonService.DeleteLesson(id)

	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func UpdateLesson(w http.ResponseWriter, r *http.Request) {
	id := common.StrToInt64(chi.URLParam(r, "id"))
	var reqBody struct {
		Title             string
		Quota             int64
		LessonTime        string
		TeacherIdOfLesson int64
	}
	err := common.BodyToJsonReq(r, &reqBody)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = services.GetServices(svc.App.DB).LessonService.UpdateLesson(models.Lesson{
		ID:                id,
		Title:             reqBody.Title,
		Quota:             reqBody.Quota,
		LessonTime:        reqBody.LessonTime,
		TeacherIdOfLesson: reqBody.TeacherIdOfLesson,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, "ok")
}

func ListLesson(w http.ResponseWriter, r *http.Request) {

	data, err := services.GetServices(svc.App.DB).LessonService.ListLesson()
	if err != nil {
		log.Fatal(err)
		return
	}
	mhttp.ResponseSuccess(w, data)
}
