package database

import (
	DataCollector "rub_buddy/features/collectors"
	DataUsers "rub_buddy/features/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(DataUsers.User{})
	db.AutoMigrate(DataCollector.Collectors{})
	return nil
}
