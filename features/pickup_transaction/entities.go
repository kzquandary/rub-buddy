package pickuptransaction

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type PickupTransaction struct {
	ID              uint          `json:"pickup_transaction_id"`
	PickupRequestID uint          `json:"pickup_request_id"`
	CollectorID     uint          `json:"collector_id"`
	TpsID           uint          `json:"tps_id"`
	PickupTime      time.Time     `json:"pickup_time"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	DeletedAt       *sql.NullTime `json:"deleted_at"`
}

type PickupTransactionCreate struct {
	ID              uint      `json:"pickup_transaction_id"`
	PickupRequestID uint      `json:"pickup_request_id"`
	CollectorID     uint      `json:"collector_id"`
	TpsID           uint      `json:"tps_id"`
	PickupTime      time.Time `json:"pickup_time"`
	ChatID          uint      `json:"chat_id"`
}

type PickupTransactionHandlerInterface interface {
	CreatePickupTransaction() echo.HandlerFunc
	GetAllPickupTransaction() echo.HandlerFunc
	GetPickupTransactionByID() echo.HandlerFunc
}

type PickupTransactionServiceInterface interface {
	CreatePickupTransaction(newData PickupTransaction) (*PickupTransactionCreate, error)
	GetAllPickupTransaction(userId uint) ([]PickupTransaction, error)
	GetPickupTransactionByID(id int) (PickupTransaction, error)
}

type PickupTransactionDataInterface interface {
	CreatePickupTransaction(newData PickupTransaction) (*PickupTransactionCreate, error)
	GetAllPickupTransaction(userId uint) ([]PickupTransaction, error)
	GetPickupTransactionByID(id int) (PickupTransaction, error)
}
