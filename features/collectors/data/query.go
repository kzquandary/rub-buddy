package data

import (
	"math/rand"
	"rub_buddy/constant"
	"rub_buddy/constant/tablename"
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
		return nil, constant.ErrCollectorUserEmailExists
	}

	newCollector.CreatedAt = time.Now()
	newCollector.UpdatedAt = time.Now()
	newCollector.ID = uint(rand.Intn(900000) + 100000)
	err = data.DB.Create(&newCollector).Error
	if err != nil {
		return nil, constant.ErrorCollectorRegister
	}
	return &newCollector, nil
}

func (data *CollectorsData) Login(email string, password string) (*collectors.Collectors, error) {
	var collector = new(Collector)
	collector.Email = email
	collector.Password = password

	var collectorData = data.DB.Where("email = ?", collector.Email).First(collector)

	var dataCount int64
	collectorData.Count(&dataCount)

	if dataCount == 0 {
		return nil, constant.ErrCollectorUserNotFound
	}
	if !helper.CheckPasswordHash(password, collector.Password) {
		return nil, constant.ErrCollectorIncorrectPassword
	}

	var result = new(collectors.Collectors)
	result.ID = collector.ID
	result.Name = collector.Name
	result.Email = collector.Email
	result.Password = collector.Password
	result.Gender = collector.Gender
	return result, nil
}

func (data *CollectorsData) UpdateCollector(collector *collectors.CollectorUpdate) error {
	var existingCollector collectors.Collectors
	err := data.DB.Table(tablename.CollectorTableName).Where("id = ?", collector.ID).First(&existingCollector).Error
	if err != nil {
		return constant.ErrCollectorUserNotFound
	}

	if collector.Email != existingCollector.Email {
		var count int64
		data.DB.Table(tablename.CollectorTableName).Where("email = ?", collector.Email).Count(&count)
		if count > 0 {
			return constant.ErrUpdateCollectorEmailExists
		}
	}

	collector.UpdatedAt = time.Now()
	err = data.DB.Table(tablename.CollectorTableName).Where("id = ?", collector.ID).Updates(collector).Error
	if err != nil {
		return constant.ErrorUpdateCollector
	}
	return nil
}

func (data *CollectorsData) GetCollectorByEmail(email string) (*collectors.Collectors, error) {
	var collector collectors.Collectors
	err := data.DB.Where("email = ?", email).First(&collector).Error
	if err != nil {
		return nil, constant.ErrCollectorUserNotFound
	}
	return &collector, nil
}
