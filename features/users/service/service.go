package service

import (
	"rub_buddy/features/users"
	"rub_buddy/helper"

	"github.com/labstack/echo/v4"
)

type UserService struct {
	d users.UserDataInterface
	j helper.JWTInterface
}

func New(data users.UserDataInterface, jwt helper.JWTInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		j: jwt,
	}
}

func (s *UserService) Register(newUser *users.User) error {
	checkUser := s.d.GetUser(newUser)

	if checkUser.Email == newUser.Email {
		return echo.NewHTTPError(400, "Email already registered")
	}
	
}
