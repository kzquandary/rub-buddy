package chatmessage

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ChatMessage struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chat_id"`
	Content   string    `json:"content"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ChatMessageHandlerInterface interface {
	CreateChatMessage() echo.HandlerFunc
}

type ChatMessageServiceInterface interface {
	CreateChatMessage(newData ChatMessage) (*ChatMessage, error)
}

type ChatMessageDataInterface interface {
	CreateChatMessage(newData *ChatMessage) (*ChatMessage, error)
}
