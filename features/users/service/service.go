package service

import (
	"rub_buddy/constant"
	"rub_buddy/features/users"
	"rub_buddy/helper"
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
	if email == "" || password == "" {
		return nil, constant.ErrLoginEmptyInput
	}

	result, err := s.d.Login(email, password)

	if err != nil {
		return nil, err
	}

	token, err := s.j.GenerateJWT(result.ID, constant.RoleUser, result.Email, result.Address)
	if err != nil {
		return nil, err
	}

	response := new(users.UserCredentials)
	response.ID = result.ID
	response.Email = result.Email
	response.Token = token
	return response, nil
}

func (s *UserService) Register(user users.User) (*users.User, error) {
	result, err := s.d.Register(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) UpdateUser(user *users.UserUpdate) error {
	return s.d.UpdateUser(user)
}

func (s *UserService) GetUserByEmail(email string) (*users.User, error) {
	return s.d.GetUserByEmail(email)
}
