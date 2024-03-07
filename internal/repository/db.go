package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDb() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "db.db")

	if err != nil {
		return nil, err
	}

	query := `
		
		CREATE TABLE IF NOT  EXISTS users(id INTEGER PRIMARY KEY, username TEXT,email TEXT NOT NULL UNIQUE,hash_password TEXT NOT NULL,role TEXT NOT NULL);
		CREATE TABLE IF NOT EXISTS sessions(id INTEGER PRIMARY KEY,user_id INTEGER,token TEXT NOT NULL UNIQUE,expired_date TEXT NOT NULL,FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE);	
		CREATE TABLE IF NOT EXISTS products(id INTEGER PRIMARY KEY,name TEXT NOT NULL, category TEXT NOT NULL, product_type TEXT, year TEXT, age_category TEXT, chrono TEXT, key_words TEXT,description TEXT,director TEXT,producer TEXT)
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfuly connected to database")
	return db, nil
}
