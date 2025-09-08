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

	createWalletTable := `
		CREATE TABLE IF NOT EXISTS wallets(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			balance DECIMAL(10, 2) DEFAULT 0.00,
			interestEarned DECIMAL(12,2) DEFAULT 0.00,
			createdAt TIMESTAMP NOT NULL,
			userId INTEGER,
			FOREIGN KEY(userId) REFERENCES users(id)
		)
	`
	_ , err = DB.Exec(createWalletTable)
	if err != nil {
		panic("Couldnot create table!")
	}
}