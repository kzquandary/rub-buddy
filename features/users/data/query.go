package data

import (
	"math/rand"
	"rub_buddy/constant"
	"rub_buddy/constant/tablename"
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
		return nil, constant.ErrRegisterUserExists
	}
	newUser.ID = uint(rand.Intn(900000) + 100000)
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
		return nil, constant.UserNotFound
	}
	if !helper.CheckPasswordHash(password, user.Password) {
		return nil, constant.ErrLoginIncorrectPassword
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

func (data *UserData) UpdateUser(user *users.UserUpdate) error {
	var existingUser User
	err := data.DB.Table(tablename.UserTableName).Where("id = ?", user.ID).First(&existingUser).Error
	if err != nil {
		return constant.UserNotFound
	}

	if user.Email != existingUser.Email {
		var count int64
		data.DB.Table(tablename.UserTableName).Where("email = ?", user.Email).Count(&count)
		if count > 0 {
			return constant.ErrUpdateUserEmailExists
		}
	}

	user.UpdatedAt = time.Now()
	err = data.DB.Table(tablename.UserTableName).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return constant.ErrUpdateUser
	}
	return nil
}

func (data *UserData) GetUserByEmail(email string) (*users.User, error) {
	var user = new(User)
	user.Email = email

	var query = data.DB.Where("email = ?", user.Email).First(user)

	var dataCount int64
	query.Count(&dataCount)

	if dataCount == 0 {
		return nil, constant.UserNotFound
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
