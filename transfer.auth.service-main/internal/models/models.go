package models

import "time"

// Пользователь
type User struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	CountryCode string    `json:"country_code,omitempty"`
	BirthDate   string    `json:"birthDate,omitempty"`
	Password    string    `json:"password,omitempty"`
	Blocked     bool      `json:"blocked"`
	Image       string    `json:"image,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
