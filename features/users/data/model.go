package data

import (
	"rub_buddy/constant/tablename"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name     string `gorm:"column:name;type:varchar(255);"`
	Email    string `gorm:"column:email;unique;type:varchar(255);"`
	Password string `gorm:"column:password;type:varchar(255);"`
	Address  string `gorm:"column:address;type:text;"`
	Gender   string `gorm:"column:gender;type:enum('Laki-laki','Perempuan');"`
	Balance  int64  `gorm:"column:balance;type:bigint;"`
}

func (User) TableName() string {
	return tablename.UserTableName
}
