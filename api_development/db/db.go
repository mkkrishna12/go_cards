package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func createTable() {
	createTables := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description STRING NOT NULL,
	location STRING NOT NULL,
	datetime DATETIME NOT NULL,
	user_id INTEGER
	)`

	if _, err := DB.Exec(createTables); err != nil {
		panic(err)
	}

}

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api_test.db")

	if err != nil {
		panic("Coudnot connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}
