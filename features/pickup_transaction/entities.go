package pickuptransaction

import (
	"time"

	"github.com/labstack/echo/v4"
)

type PickupTransaction struct {
	ID              uint      `json:"pickup_transaction_id"`
	PickupRequestID uint      `json:"pickup_request_id"`
	CollectorID     uint      `json:"collector_id"`
	TpsID           uint      `json:"tps_id"`
	PickupTime      time.Time `json:"pickup_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type PickupTransactionHandlerInterface interface {
	CreatePickupTransaction() echo.HandlerFunc
	GetAllPickupTransaction() echo.HandlerFunc
	GetPickupTransactionByID() echo.HandlerFunc
	DeletePickupTransactionByID() echo.HandlerFunc
}

type PickupTransactionServiceInterface interface {
	CreatePickupTransaction(newData PickupTransaction) (*PickupTransaction, error)
	GetAllPickupTransaction() ([]PickupTransaction, error)
	GetPickupTransactionByID(PickupTransaction) (PickupTransaction, error)
	DeletePickupTransactionByID(PickupTransaction) error
}

type PickupTransactionDataInterface interface {
	CreatePickupTransaction(newData PickupTransaction) (*PickupTransaction, error)
	GetAllPickupTransaction() ([]PickupTransaction, error)
	GetPickupTransactionByID(PickupTransaction) (PickupTransaction, error)
	DeletePickupTransactionByID(PickupTransaction) error
}
