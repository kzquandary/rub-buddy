package pickuprequest

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type PickupRequest struct {
	ID          uint          `json:"pickup_request_id"`
	UserID      uint          `json:"user_id"`
	Weight      float64       `json:"weight"`
	Address     string        `json:"address"`
	Description string        `json:"description"`
	Earnings    float64       `json:"earnings"`
	Image       string        `json:"image"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   *sql.NullTime `json:"deleted_at"`
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
	DeletePickupRequestByID(id int, userID uint) error
}

type PickupRequestDataInterface interface {
	CreatePickupRequest(newData *PickupRequest) (*PickupRequest, error)
	GetAllPickupRequest() ([]PickupRequest, error)
	GetPickupRequestByID(id int) (PickupRequest, error)
	DeletePickupRequestByID(id int, userID uint) error
}
