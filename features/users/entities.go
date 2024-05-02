package users

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newUser *User) error
	Login(email string, password string) error
	GetUser(user *User)
	UpdateUser(user *User) error
}

type UserDataInterface interface {
	Register(user *User) error
	Login(user *User) error
	GetUser(user *User) error
	UpdateUser(user *User) error
}
