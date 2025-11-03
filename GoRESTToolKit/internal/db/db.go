package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(connStr string) error {
	var err error
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	log.Println("Connected to the database")
	return nil
}
