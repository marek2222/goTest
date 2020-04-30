package test

import (
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

var db *sqlx.DB

// Start uje serwis samoloty-miejsca
func Start(ruter *httprouter.Router, dbx *sqlx.DB) {
	db = dbx
	ruter.GET("/samoloty-miejsca/:samolotid", odczytHandler)
	ruter.POST("/samoloty-miejsca/:samolotid", zapisHandler)
	ruter.DELETE("/samoloty-miejsca/:samolotid", usuwanieHandler)
}
