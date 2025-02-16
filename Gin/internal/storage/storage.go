package storage

import (
	"database/sql"
	"errors"
	"gin-api/internal/models"
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

	// Create table if not exists
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

func (s *EmployeeStorage) GetEmployees() ([]models.Employee, error) {
	rows, err := s.db.Query("SELECT ID, Type, Username, Password FROM Employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.ID, &emp.Type, &emp.Username, &emp.Password)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func (s *EmployeeStorage) GetEmployee(id int) (*models.Employee, error) {
	var emp models.Employee
	err := s.db.QueryRow("SELECT ID, Type, Username, Password FROM Employee WHERE ID = ?", id).Scan(
		&emp.ID, &emp.Type, &emp.Username, &emp.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEmployeeNotFound
		}
		return nil, err
	}
	return &emp, nil
}

func (s *EmployeeStorage) CreateEmployee(emp *models.Employee) error {
	res, err := s.db.Exec("INSERT INTO Employee (Type, Username, Password) VALUES (?, ?, ?)",
		emp.Type, emp.Username, emp.Password)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	emp.ID = int(id)
	return nil
}

func (s *EmployeeStorage) UpdateEmployee(id int, emp *models.Employee) error {
	_, err := s.db.Exec("UPDATE Employee SET Type = ?, Username = ?, Password = ? WHERE ID = ?",
		emp.Type, emp.Username, emp.Password, id)
	return err
}

func (s *EmployeeStorage) DeleteEmployee(id int) error {
	_, err := s.db.Exec("DELETE FROM Employee WHERE ID = ?", id)
	return err
}
