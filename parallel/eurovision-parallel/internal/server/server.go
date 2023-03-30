package server

import (
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
)

var App struct {
	DB    *sqlx.DB
	Store *sessions.CookieStore
}
