package data

import "gorm.io/gorm"

type PickupTransaction struct {
	*gorm.Model
	ID          uint    `gorm:"primaryKey;column:id;autoIncrement;not null;type:varchar(255)"`
	UserID      uint    `gorm:"column:user_id;not null;type:varchar(255)"`
	CollectorID string  `gorm:"column:collector_id;not null;type:varchar(255)"`
	PickupTime  string  `gorm:"column:pickup_time;not null;type:varchar(255)"`
	TpsID       uint    `gorm:"column:tps_id;not null;type:varchar(255)"`
}

func (PickupTransaction) TableName() string {
	return "pickup_transaction"
}
