package entities

type User struct {
	ID       int64
	Name     string
	Username string
	Email    string
}

func NewUser(name, username, email, password string) User {
	return User{
		Name:     name,
		Username: username,
		Email:    email,
	}
}
