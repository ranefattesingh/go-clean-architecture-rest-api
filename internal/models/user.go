package models

import (
	"time"
)

type User struct {
	UserID      int        `json:"user_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email,omitempty"`
	Password    string     `json:"password,omitempty"`
	Role        *string    `json:"role,omitempty"`
	About       *string    `json:"about,omitempty"`
	Avatar      *string    `json:"avatar,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Address     *string    `json:"address,omitempty"`
	City        *string    `json:"city,omitempty"`
	Country     *string    `json:"country,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Postcode    *int       `json:"postcode,omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	LoginDate   time.Time  `json:"login_date"`
}
