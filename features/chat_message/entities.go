package chatmessage

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type ChatMessage struct {
	ID        uint          `json:"id"`
	ChatID    uint          `json:"chat_id"`
	Content   string        `json:"content"`
	Sender    uint          `json:"sender"`
	Receiver  uint          `json:"receiver"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *sql.NullTime `json:"deleted_at"`
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
