package service

import (
	"errors"
	"rub_buddy/constant"
	"rub_buddy/features/collectors"
	"rub_buddy/helper"
	"strings"
	"time"
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
	_, err := s.d.GetCollectorByEmail(newCollector.Email)
	if err == nil {
		return nil, errors.New(constant.EmailAlreadyExists)
	}

	result, err := s.d.Register(newCollector)
	if err != nil {
		return nil, errors.New("Internal server error")
	}
	return result, nil
}

func (s *CollectorService) Login(email string, password string) (*collectors.CollectorCredentials, error) {
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

	token, err := s.j.GenerateJWT(result.ID, "Collector", result.Email, "")
	if err != nil {
		return nil, errors.New("Internal server error")
	}

	response := new(collectors.CollectorCredentials)
	response.ID = result.ID
	response.Email = result.Email
	response.Token = token
	return response, nil
}

func (s *CollectorService) UpdateCollector(collector *collectors.Collectors) error {
	collector.UpdatedAt = time.Now()
	return s.d.UpdateCollector(collector)
}

func (s *CollectorService) GetCollector(collector *collectors.Collectors) error {
	return s.d.GetCollector(collector)
}

func (s *CollectorService) GetCollectorByEmail(email string) (*collectors.Collectors, error) {
	return s.d.GetCollectorByEmail(email)
}
