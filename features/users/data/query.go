package data

import (
	"errors"
	"rub_buddy/constant"
	"rub_buddy/features/users"
	"rub_buddy/helper"
	"time"

	"gorm.io/gorm"
)

type UserData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		DB: db,
	}
}

func (data *UserData) Register(newUser users.User) (*users.User, error) {
	_, err := data.GetUserByEmail(newUser.Email)
	if err == nil {
		return nil, errors.New(constant.EmailAlreadyExists)
	}
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	return &newUser, data.DB.Create(&newUser).Error
}

func (data *UserData) Login(email string, password string) (*users.User, error) {
	var user = new(User)
	user.Email = email
	user.Password = password

	var UserData = data.DB.Where("email = ?", user.Email).First(user)
	var userCount int64
	UserData.Count(&userCount)

	if userCount == 0 {
		return nil, errors.New(constant.NotFound)
	}
	if !helper.CheckPasswordHash(password, user.Password) {
		return nil, errors.New(constant.IncorrectPassword)
	}

	var result = new(users.User)
	result.ID = user.ID
	result.Name = user.Name
	result.Email = user.Email
	result.Password = user.Password
	result.Address = user.Address
	result.Gender = user.Gender
	return result, nil
}

func (data *UserData) GetUser(user *users.User) error {
	return data.DB.Where("email = ?", user.Email).First(user).Error
}

func (data *UserData) UpdateUser(user *users.User) error {
	user.UpdatedAt = time.Now()
	return data.DB.Save(&user).Error
}

func (data *UserData) GetUserByEmail(email string) (*users.User, error) {
	var user = new(User)
	user.Email = email

	var query = data.DB.Where("email = ?", user.Email).First(user)

	var dataCount int64
	query.Count(&dataCount)

	if dataCount == 0 {
		return nil, errors.New(constant.NotFound)
	}

	var result = new(users.User)
	result.ID = user.ID
	result.Name = user.Name
	result.Email = user.Email
	result.Password = user.Password
	result.Address = user.Address
	result.Gender = user.Gender
	result.CreatedAt = user.CreatedAt
	result.UpdatedAt = user.UpdatedAt
	return result, nil
}
