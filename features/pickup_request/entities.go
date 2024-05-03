package pickuprequest

import (
	"time"

	"github.com/labstack/echo/v4"
)

type PickupRequest struct {
	ID          uint      `json:"pickup_request_id"`
	UserID      uint      `json:"user_id"`
	Weight      int       `json:"weight"`
	Description string    `json:"description"`
	Earnings    int       `json:"earnings"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PickupRequestHandlerInterface interface {
	CreatePickupRequest() echo.HandlerFunc
	GetAllPickupRequest() echo.HandlerFunc
	GetPickupRequestByID() echo.HandlerFunc
	DeletePickupRequestByID() echo.HandlerFunc
}

type PickupRequestServiceInterface interface {
	CreatePickupRequest(newData PickupRequest) (*PickupRequest, error)
	GetAllPickupRequest() ([]PickupRequest, error)
	GetPickupRequestByID(id int) (PickupRequest, error)
	DeletePickupRequestByID(id int) error
}

type PickupRequestDataInterface interface {
	CreatePickupRequest(newData *PickupRequest) (*PickupRequest, error)
	GetAllPickupRequest() ([]PickupRequest, error)
	GetPickupRequestByID(id int) (PickupRequest, error)
	DeletePickupRequestByID(id int) error
}
