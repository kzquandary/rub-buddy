package data

import "gorm.io/gorm"

type Payment struct {
	*gorm.Model
	OrderID string `gorm:"column:order_id;not null;primaryKey;type:varchar(255)"`
	UserID  uint   `gorm:"column:user_id;not null;type:varchar(255)"`
	Amount  int    `gorm:"column:amount;not null;type:varchar(255)"`
	Status  string `gorm:"column:status;not null;type:varchar(255)"`
	SnapURL string `gorm:"column:snap_url;not null;type:varchar(255)"`
}

func (m *Payment) TableName() string {
	return "payments"
}
