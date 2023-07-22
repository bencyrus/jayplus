package models

import "time"

type User struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        string    `json:"role"` // Possible values: "customer", "employee", "admin", "owner"
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
