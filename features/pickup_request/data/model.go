package data

import "gorm.io/gorm"

type PickupRequest struct {
	*gorm.Model
	ID          uint    `gorm:"primaryKey;column:id;autoIncrement;not null;type:varchar(255)"`
	UserID      uint    `gorm:"column:user_id;not null;type:varchar(255)"`
	Weight      float64 `gorm:"column:weight;not null;type:varchar(255)"`
	Address     string  `gorm:"column:address;not null;type:varchar(255)"`
	Description string  `gorm:"column:description;not null;type:varchar(255)"`
	Earnings    float64 `gorm:"column:earnings;not null;type:varchar(255)"`
	Image       string  `gorm:"column:image;not null;type:varchar(255)"`
	Status      string  `gorm:"column:status;not null;type:enum('Pending','Accepted','Rejected')"`
}

func (PickupRequest) TableName() string {
	return "pickup_request"
}
