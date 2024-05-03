package service

import (
	"errors"
	"rub_buddy/constant"
	"rub_buddy/features/users"
	"rub_buddy/helper"
	"strings"
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

func (s *UserService) Login(email string, password string) (*users.UserCredentials, error) {
	result, err := s.d.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), constant.NotFound) {
			return nil, errors.New("User not found")
		}
		if strings.Contains(err.Error(), constant.IncorrectPassword) {
			return nil, errors.New("Incorrect password")
		}
		return nil, errors.New("Internal server error")
	}

	token, err := s.j.GenerateJWT(result.ID, "User", result.Email)
	if err != nil {
		return nil, errors.New("Internal server error")
	}

	response := new(users.UserCredentials)
	response.ID = result.ID
	response.Email = result.Email
	response.Token = token
	return response, nil
}

func (s *UserService) Register(user users.User) (*users.User, error) {
	_, err := s.d.GetUserByEmail(user.Email)
	if err == nil {
		return nil, errors.New("User already exists")
	}

	result, err := s.d.Register(user)
	if err != nil {
		return nil, errors.New("Internal server error")
	}
	return result, nil
}

func (s *UserService) UpdateUser(user *users.User) error {
	return s.d.UpdateUser(user)
}

func (s *UserService) GetUser(user *users.User) error {
	return s.d.GetUser(user)
}

func (s *UserService) GetUserByEmail(email string) (*users.User, error) {
	return s.d.GetUserByEmail(email)
}
