package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	dsn := "golang_user:golang_password@tcp(localhost:3306)/golang_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}
