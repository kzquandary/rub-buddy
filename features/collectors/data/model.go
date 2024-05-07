package data

import (
	"gorm.io/gorm"

	"rub_buddy/constant/tablename"
)

type Collector struct {
	*gorm.Model
	Name     string `gorm:"column:name;type:varchar(255);"`
	Email    string `gorm:"column:email;unique;type:varchar(255);"`
	Password string `gorm:"column:password;type:varchar(255);"`
	Gender   string `gorm:"column:gender;type:enum('Laki-laki','Perempuan');"`
}

func (Collector) TableName() string {
	return tablename.CollectorTableName
}
