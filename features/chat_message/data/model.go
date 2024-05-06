package data

import "gorm.io/gorm"

type ChatMessage struct {
	*gorm.Model
	ChatID   uint   `gorm:"column:chat_id;not null;type:varchar(255)"`
	Content  string `gorm:"column:content;not null;type:varchar(255)"`
	Sender   uint   `gorm:"column:user_id;not null;type:varchar(255)"`
	Receiver uint   `gorm:"column:collector_id;not null;type:varchar(255)"`
}

func (ChatMessage) TableName() string {
	return "chat_messages"
}
