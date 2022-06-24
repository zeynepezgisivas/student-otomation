package models

type SetGrades struct {
	StudentNo     int64  `db:"student_no"`
	TitleOfLesson string `db:"title_of_lesson"`
	TypeOfNote    string `db:"type_of_note"`
	ValueOfNote   int64  `db:"value_of_note"`
}

type AverageReq struct {
	ID           int64  `db:"id"`
	LessonID     int64  `db:"lesson_id"`
	SchoolNumber int64  `db:"school_number"`
	LessonTitle  string `db:"title_of_lesson"`
}
