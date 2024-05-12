package pickuprequest

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type PickupRequest struct {
	ID          uint          
	UserID      uint          
	Weight      float64       
	Address     string        
	Description string        
	Earnings    float64       
	Image       string        
	Status      string        
	CreatedAt   time.Time     
	UpdatedAt   time.Time     
	DeletedAt   *sql.NullTime 
}

type PickupRequestInfo struct {
	ID          uint

	UserID      uint
	UserName    string

	Weight      float64
	Address     string
	Description string
	Earnings    float64
	Image       string
	Status      string
}

type PickupRequestHandlerInterface interface {
	CreatePickupRequest() echo.HandlerFunc
	GetAllPickupRequest() echo.HandlerFunc
	GetPickupRequestByID() echo.HandlerFunc
	DeletePickupRequestByID() echo.HandlerFunc
}

type PickupRequestServiceInterface interface {
	CreatePickupRequest(newData PickupRequest) (*PickupRequest, error)
	GetAllPickupRequest(role string, id uint) ([]PickupRequestInfo, error)
	GetPickupRequestByID(id int) (PickupRequestInfo, error)
	DeletePickupRequestByID(id int, userID uint) error
}

type PickupRequestDataInterface interface {
	CreatePickupRequest(newData *PickupRequest) (*PickupRequest, error)
	GetAllPickupRequest(role string, id uint) ([]PickupRequestInfo, error)
	GetPickupRequestByID(id int) (PickupRequestInfo, error)
	DeletePickupRequestByID(id int, userID uint) error
}
