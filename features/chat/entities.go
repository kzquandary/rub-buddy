package chat

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type Chat struct {
	ID                  uint          `json:"chat_id"`
	PickupTransactionID uint          `json:"pickup_transaction_id"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	DeletedAt           *sql.NullTime `json:"deleted_at"`
}

type ChatHandlerInterface interface {
	Chat() echo.HandlerFunc
	GetChatByID() echo.HandlerFunc
}

type ChatServiceInterface interface {
	CreateChat(chat *Chat) (*Chat, error)
	GetChatByID(chat *Chat) (*Chat, error)
}

type ChatDataInterface interface {
	CreateChat(chat *Chat) (*Chat, error)
	GetChatByID(chat *Chat) (*Chat, error)
}
