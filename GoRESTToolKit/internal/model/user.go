package model

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"-"` // Never expose password in JSON
	Roles    []string `json:"roles"`
}
