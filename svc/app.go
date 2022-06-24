package svc

import (
	"github.com/jmoiron/sqlx"
)

var App *app

type app struct {
	DB *sqlx.DB
}

func InitApp(db *sqlx.DB) {
	App = &app{
		DB: db,
	}
}
