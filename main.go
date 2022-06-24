package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"studentOtomasyon/api"
	"studentOtomasyon/svc"
)

func main() {
	//İlk önce database bağlantısı yapalım
	psqlInfo := "host=localhost port=5432 user=postgres password=tayitkan dbname=student_otomasyon"
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	svc.InitApp(db)

	router := chi.NewRouter()

	//lesson endpoints
	router.Post("/lesson/create", api.CreateLesson)
	router.Get("/lesson/{id}", api.GetLesson)
	router.Delete("/lesson/{id}", api.DeleteLesson)
	router.Put("/lesson/{id}", api.UpdateLesson)
	router.Get("/lesson", api.ListLesson)

	router.Post("/student/create", api.CreateStudent)
	router.Get("/student/{id}", api.GetStudent)
	router.Delete("/student/{id}", api.DeleteStudent)
	router.Put("/student/{id}", api.UpdateStudent)
	router.Get("/student", api.ListStudent)

	router.Post("/teacher/create", api.CreateTeacher)
	router.Get("/teacher/{id}", api.GetTeacher)
	router.Delete("/teacher/{id}", api.DeleteTeacher)
	router.Put("/teacher/{id}", api.UpdateTeacher)
	router.Get("/teacher", api.ListTeacher)

	router.Post("/testapi", api.Testapi)

	log.Println("listening on", 18000)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 18000), router))
}

/*
Apiler:
	1. Generic Apiler:
		1.1. Student table:
			1.1.1. Create
			1.1.2. Get
			1.1.3. Delete
			1.1.4. Update
			1.1.5. List
		1.2. Teacher table:
			1.2.1. Create
			1.2.2. Get
			1.2.3. Delete
			1.2.4. Update
			1.2.5. List
		1.3. Lesson table:
			1.3.1. Create
			1.3.2. Get
			1.3.3. Delete
			1.3.4. Update
			1.3.5. List
	2. Özel Apiler:
		2.1. Grades Apileri
			2.1.1. Öğrenciye not girme
			2.1.2  Öğrencinin notlarını getir.
			2.1.3  Öğrencinin ortalamasını getir.
			2.1.4  Hocanın veya dersin notlarını getir.
			2.1.5  Dersin veya hocanın not ortalamasını getir.
		2.2. Lesson_student Apileri
			2.2.1  Derse öğrenci ekleme
			2.2.2  Derten öğrenci çıkarma
		2.3. Teacher Apileri
			2.3.1  Öğretmen verdiği dersleri görsün.
*/
