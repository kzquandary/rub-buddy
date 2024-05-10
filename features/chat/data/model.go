package data

import (
	"rub_buddy/constant/tablename"

	"gorm.io/gorm"
)

type Chat struct {
	*gorm.Model
	ID                  uint `gorm:"primaryKey;column:id;autoIncrement;not null;type:varchar(255)"`
	PickupTransactionID uint `gorm:"column:pickup_transaction_id;not null;type:varchar(255)"`
}

func (Chat) TableName() string {
	return tablename.ChatTableName
}

