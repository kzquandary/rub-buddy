package midtranspayment

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Midtrans struct {
	ID        string
	Amount    int64
	UserID    uint
	Status    int
	SnapURL   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type MidtransHandlerInterface interface {
	CreateTransaction() echo.HandlerFunc
	VerifyPayment() echo.HandlerFunc
}

type MidtransServiceInterface interface {
	GenerateSnapURL(payment Midtrans) (Midtrans, error)
	VerifyPayment(orderId string) error
}

type MidtransDataInterface interface {
	CreatePayment(payment Midtrans) error
	VerifyPayment(orderId string) error
}
