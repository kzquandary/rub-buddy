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

type ChatInfo struct {
	ID                  uint
	PickupTransactionID uint

	UserID        uint
	UserName      string
	CollectorID   uint
	CollectorName string
}
type ChatHandlerInterface interface {
	GetChat() echo.HandlerFunc
}

type ChatServiceInterface interface {
	GetChat(id uint, role string) ([]ChatInfo, error)
}

type ChatDataInterface interface {
	GetChat(id uint, role string) ([]ChatInfo, error)
}
