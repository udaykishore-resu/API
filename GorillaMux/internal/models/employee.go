package models

type Employee struct {
	ID       int    `json:"id" db:"id"`
	Type     string `json:"type" db:"type"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
