package storage

import (
	"database/sql"
	"errors"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeStorage struct {
	db *sql.DB
}

func NewEmployeeStorage(dsn string) (*EmployeeStorage, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Employee (
        ID INT PRIMARY KEY AUTO_INCREMENT,
        Type VARCHAR(255),
        Username VARCHAR(255),
        Password VARCHAR(255)
    )`)

	if err != nil {
		return nil, err
	}

	return &EmployeeStorage{db: db}, nil

}
