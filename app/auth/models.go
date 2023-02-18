package auth

type User struct {
	Email      string
	FirstName  string
	SecondName string
	ThirdName  string
}

type SignIn struct {
	Email    string
	Password string
}
