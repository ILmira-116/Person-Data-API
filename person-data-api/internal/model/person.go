package model

import "time"

// Person модель человека
type Person struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Surname     string    `json:"surname" db:"surname"`
	Patronymic  string    `json:"patronymic,omitempty" db:"patronymic"`
	Age         int       `json:"age,omitempty" db:"age"`
	Gender      string    `json:"gender,omitempty" db:"gender"`
	Nationality string    `json:"nationality,omitempty" db:"nationality"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
