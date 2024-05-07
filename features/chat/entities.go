package chat

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type Chat struct {
	ID                  uint          
	PickupTransactionID uint          
	CreatedAt           time.Time     
	UpdatedAt           time.Time     
	DeletedAt           *sql.NullTime 
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
