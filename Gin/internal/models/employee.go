package models

type Employee struct {
	ID       int    `json:"id" db:"id"`
	Type     string `json:"type" db:"type" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"passwword" binding:"required"`
}
