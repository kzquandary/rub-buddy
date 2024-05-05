package database

import (
	"gorm.io/gorm"
)

func Truncate(db *gorm.DB) error {
	db.Exec("TRUNCATE TABLE users CASCADE")
	db.Exec("TRUNCATE TABLE collectors CASCADE")
	db.Exec("TRUNCATE TABLE pickup_requests CASCADE")
	db.Exec("TRUNCATE TABLE pickup_transactions CASCADE")
	db.Exec("TRUNCATE TABLE chats CASCADE")
	db.Exec("TRUNCATE TABLE chat_messages CASCADE")
	return nil
}
