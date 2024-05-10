package service

import (
	"rub_buddy/constant"
	"rub_buddy/features/collectors"
	"rub_buddy/helper"
)

type CollectorService struct {
	d collectors.CollectorDataInterface
	j helper.JWTInterface
}

func New(data collectors.CollectorDataInterface, jwt helper.JWTInterface) collectors.CollectorServiceInterface {
	return &CollectorService{
		d: data,
		j: jwt,
	}
}

func (s *CollectorService) Register(newCollector collectors.Collectors) (*collectors.Collectors, error) {
	if newCollector.Email == "" || newCollector.Password == "" || newCollector.Name == "" || newCollector.Gender == "" {
		return nil, constant.ErrRegisterEmptyInput
	}
	result, err := s.d.Register(newCollector)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *CollectorService) Login(email string, password string) (*collectors.CollectorCredentials, error) {
	if email == "" || password == "" {
		return nil, constant.ErrLoginEmptyInput
	}
	result, err := s.d.Login(email, password)

	if err != nil {
		return nil, err
	}

	token, err := s.j.GenerateJWT(result.ID, constant.RoleCollector, result.Email, "")
	if err != nil {
		return nil, err
	}

	response := new(collectors.CollectorCredentials)
	response.ID = result.ID
	response.Email = result.Email
	response.Token = token
	return response, nil
}

func (s *CollectorService) UpdateCollector(collector *collectors.CollectorUpdate) error {
	return s.d.UpdateCollector(collector)
}

func (s *CollectorService) GetCollectorByEmail(email string) (*collectors.Collectors, error) {
	return s.d.GetCollectorByEmail(email)
}
