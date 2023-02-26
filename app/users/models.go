package users

import "time"

type User struct {
	ID             string    `json:"ID"`
	Email          string    `json:"email" validate:"required,email"`
	PhoneNumber    string    `json:"phoneNumber" validate:"required,email"`
	FirstName      string    `json:"firstName"`
	SecondName     string    `json:"secondName"`
	ThirdName      string    `json:"thirdName"`
	EmailConfirmed bool      `json:"-"`
	RegisterTime   time.Time `json:"-"`
}

type SingUp struct {
	User
	Password string `json:"password"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
