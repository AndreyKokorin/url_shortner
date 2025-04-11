package model

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserRepository interface {
	NewUser(user *User) error
	GetByEmail(email string) (*User, error)
}

type UserService interface {
	RegisterUser(user *User) error
	LogIn(email string, password string) (string, error)
}
