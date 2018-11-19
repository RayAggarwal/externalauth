package models

import (
	"database/sql"
)

var db* sql.DB

func InitDb(datasourceName) {
	db, err := sql.Open("mysql", "")
}