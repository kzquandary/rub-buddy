package data

import (
	"errors"
	"rub_buddy/constant"
	"rub_buddy/features/collectors"
	"rub_buddy/helper"
	"time"

	"gorm.io/gorm"
)

type CollectorsData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) collectors.CollectorDataInterface {
	return &CollectorsData{
		DB: db,
	}
}

func (data *CollectorsData) Register(newCollector collectors.Collectors) (*collectors.Collectors, error) {
	_, err := data.GetCollectorByEmail(newCollector.Email)
	if err == nil {
		return nil, errors.New(constant.EmailAlreadyExists)
	}

	newCollector.CreatedAt = time.Now()
	newCollector.UpdatedAt = time.Now()

	return &newCollector, data.DB.Create(&newCollector).Error
}

func (data *CollectorsData) Login(email string, password string) (*collectors.Collectors, error) {
	var collector = new(Collector)
	collector.Email = email
	collector.Password = password

	var collectorData = data.DB.Where("email = ?", collector.Email).First(collector)

	var dataCount int64
	collectorData.Count(&dataCount)

	if dataCount == 0 {
		return nil, errors.New(constant.NotFound)
	}
	if !helper.CheckPasswordHash(password, collector.Password) {
		return nil, errors.New(constant.IncorrectPassword)
	}

	var result = new(collectors.Collectors)
	result.ID = collector.ID
	result.Name = collector.Name
	result.Email = collector.Email
	result.Password = collector.Password
	result.Gender = collector.Gender
	return result, nil
}

func (data *CollectorsData) GetCollector(collector *collectors.Collectors) error {
	result := data.DB.Where("id = ?", collector.ID).First(&collector)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (data *CollectorsData) UpdateCollector(collector *collectors.Collectors) error {
	_, err := data.GetCollectorByEmail(collector.Email)
	if err == nil {
		return errors.New(constant.EmailAlreadyExists)
	}
	result := data.DB.Where("id = ?", collector.ID).Updates(&collector)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (data *CollectorsData) GetCollectorByEmail(email string) (*collectors.Collectors, error) {
	var collector collectors.Collectors
	result := data.DB.Where("email = ?", email).First(&collector)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collector, nil
}
