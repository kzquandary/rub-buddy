package data

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name     string `gorm:"column:name;type:varchar(255);"`
	Email    string `gorm:"column:email;type:varchar(255);"`
	Password string `gorm:"column:password;type:varchar(255);"`
	Address  string `gorm:"column:address;type:text;"`
	Gender   string `gorm:"column:gender;type:enum('Laki-laki','Perempuan');"`
}

func (User) TableName() string {
	return "users"
}
