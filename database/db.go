package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "sonata.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	createTablesQuery := `
		CREATE TABLE IF NOT EXISTS music (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			filepath TEXT,
			artist TEXT
		);
		
		CREATE TABLE IF NOT EXISTS playlist (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT
		);
	`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println("Database created")
	return db, nil
}
