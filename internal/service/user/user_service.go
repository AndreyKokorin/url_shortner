package service

import (
	"URL_shortner/internal/model"
	"URL_shortner/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"os"
)

type UserService struct {
	rep model.UserRepository
}

func NewUserService(rep model.UserRepository) *UserService {
	return &UserService{rep: rep}
}

func (userService *UserService) RegisterUser(user *model.User) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Info(err.Error())
	}

	user.Password = string(hashPassword)

	err = userService.rep.NewUser(user)
	if err != nil {
		slog.Info(err.Error())
		return err
	}
	return nil
}

func (userService *UserService) LogIn(email string, password string) (string, error) {
	user, err := userService.rep.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateJWT(user.Id, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return token, nil
}
