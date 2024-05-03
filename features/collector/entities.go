package collector

import (
	"github.com/labstack/echo/v4"
)

type Collector struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetCollector() echo.HandlerFunc
	UpdateCollector() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newUser *Collector) error
	Login(email string, password string) error
	GetCollector(user *Collector)
	UpdateCollector(user *Collector) error
}

type UserDataInterface interface {
	Register(user *Collector) error
	Login(user *Collector) error
	GetCollector(user *Collector) error
	UpdateCollector(user *Collector) error
}
