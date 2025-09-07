package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Couldnot connect to database!")
	}
	
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables(){
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			firstName TEXT NOT NULL,
			lastName TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			createdAt TIMESTAMP NOT NULL
		)
	`

	_ , err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Couldnot create table!")
	}
}