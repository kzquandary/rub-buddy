package database

import (
	DataChat "rub_buddy/features/chat"
	DataChatMessage "rub_buddy/features/chat_message"
	DataCollector "rub_buddy/features/collectors"
	DataPayment "rub_buddy/features/midtranspayment"
	DataPickup "rub_buddy/features/pickup_request"
	DataPickupTransaction "rub_buddy/features/pickup_transaction"
	DataUsers "rub_buddy/features/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(DataUsers.User{})
	db.AutoMigrate(DataCollector.Collectors{})
	db.AutoMigrate(DataPickup.PickupRequest{})
	db.AutoMigrate(DataPickupTransaction.PickupTransaction{})
	db.AutoMigrate(DataChat.Chat{})
	db.AutoMigrate(DataChatMessage.ChatMessage{})
	db.AutoMigrate(DataPayment.Midtrans{})
	return nil
}
