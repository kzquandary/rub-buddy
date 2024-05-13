package pickuptransaction

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type PickupTransaction struct {
	ID              uint
	PickupRequestID uint
	CollectorID     uint
	TpsID           uint
	PickupTime      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *sql.NullTime
}

type PickupTransactionCreate struct {
	ID              uint
	PickupRequestID uint
	CollectorID     uint
	TpsID           uint
	PickupTime      time.Time
	ChatID          uint
}

type PickupTransactionInfo struct {
	ID uint

	UserID      uint
	UserName    string
	UserAddress string

	CollectorID   uint
	CollectorName string
	PickupTime    time.Time

	TpsID uint
	TpsName string
	TpsAddress string
}
type PickupTransactionHandlerInterface interface {
	CreatePickupTransaction() echo.HandlerFunc
	GetAllPickupTransaction() echo.HandlerFunc
	GetPickupTransactionByID() echo.HandlerFunc
}

type PickupTransactionServiceInterface interface {
	CreatePickupTransaction(newData PickupTransaction) (*PickupTransactionCreate, error)
	GetAllPickupTransaction(userId uint) ([]PickupTransactionInfo, error)
	GetPickupTransactionByID(id int) (PickupTransactionInfo, error)
}

type PickupTransactionDataInterface interface {
	CreatePickupTransaction(newData PickupTransaction) (*PickupTransactionCreate, error)
	GetAllPickupTransaction(userId uint) ([]PickupTransactionInfo, error)
	GetPickupTransactionByID(id int) (PickupTransactionInfo, error)
}
