package model

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserReponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
