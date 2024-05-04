package database

import (
	DataCollector "rub_buddy/features/collectors"
	DataPickup "rub_buddy/features/pickup_request"
	DataUsers "rub_buddy/features/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(DataUsers.User{})
	db.AutoMigrate(DataCollector.Collectors{})
	db.AutoMigrate(DataPickup.PickupRequest{})
	return nil
}
