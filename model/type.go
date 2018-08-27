package model

import "time"

// Model 標準のmodel
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// User 構造体
type User struct {
	Model
	Name 	string 	`json:"name"`
	Age 	int		`json:"age"`
}

type Users []User

