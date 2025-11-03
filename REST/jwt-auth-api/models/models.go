package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	Role      string `gorm:"default:'user'"`
	CreatedAt time.Time
}

type RefreshToken struct {
	ID        uint   `gorm:"primaryKey"`
	Token     string `gorm:"unique"`
	UserID    uint
	ExpiresAt time.Time
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
