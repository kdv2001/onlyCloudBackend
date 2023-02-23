package users

import "time"

type User struct {
	ID             string    `json:"ID"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phoneNumber"`
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
	Email    string
	Password string
}
