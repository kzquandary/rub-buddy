package data

import (
	"rub_buddy/constant"
	"rub_buddy/constant/tablename"
	rubbuddychat "rub_buddy/features/chat"

	"gorm.io/gorm"
)

type ChatData struct {
	*gorm.DB
}

func New(db *gorm.DB) rubbuddychat.ChatDataInterface {
	return &ChatData{
		DB: db,
	}
}

func (data *ChatData) GetChat(id uint, role string) ([]rubbuddychat.ChatInfo, error) {
	var chats []rubbuddychat.ChatInfo
	var err error

	if role == "User" {
		err = data.DB.
			Table(tablename.ChatTableName).
			Select("chats.id, chats.pickup_transaction_id, pickup_requests.user_id, users.name as user_name, pickup_transactions.collector_id, collectors.name as collector_name").
			Joins("JOIN pickup_transactions ON pickup_transactions.id = chats.pickup_transaction_id").
			Joins("JOIN pickup_requests ON pickup_requests.id = pickup_transactions.pickup_request_id").
			Joins("JOIN users ON users.id = pickup_requests.user_id").
			Joins("JOIN collectors ON collectors.id = pickup_transactions.collector_id").
			Where("users.id = ?", id).
			Scan(&chats).
			Error
	} else if role == "Collector" {
		err = data.DB.
			Table(tablename.ChatTableName).
			Select("chats.id, chats.pickup_transaction_id, pickup_requests.user_id, users.name as user_name, pickup_transactions.collector_id, collectors.name as collector_name").
			Joins("JOIN pickup_transactions ON pickup_transactions.id = chats.pickup_transaction_id").
			Joins("JOIN pickup_requests ON pickup_requests.id = pickup_transactions.pickup_request_id").
			Joins("JOIN users ON users.id = pickup_requests.user_id").
			Joins("JOIN collectors ON collectors.id = pickup_transactions.collector_id").
			Where("pickup_transactions.collector_id = ?", id).
			Scan(&chats).
			Error
	} else {
		return nil, constant.ErrChatGet
	}

	if err != nil {
		return nil, constant.ErrChatGet
	}

	return chats, nil
}
