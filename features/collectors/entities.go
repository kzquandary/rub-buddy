package collectors

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type Collectors struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *sql.NullTime
}

type CollectorCredentials struct {
	ID    uint
	Email string
	Token any
}

type CollectorUpdate struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Gender    string
	UpdatedAt time.Time
}
type CollectorHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetCollector() echo.HandlerFunc
	UpdateCollector() echo.HandlerFunc
}

type CollectorServiceInterface interface {
	Register(newUser Collectors) (*Collectors, error)
	Login(email string, password string) (*CollectorCredentials, error)
	GetCollectorByEmail(email string) (*Collectors, error)
	UpdateCollector(user *CollectorUpdate) error
}

type CollectorDataInterface interface {
	Register(collector Collectors) (*Collectors, error)
	Login(email string, password string) (*Collectors, error)
	UpdateCollector(collector *CollectorUpdate) error
	GetCollectorByEmail(email string) (*Collectors, error)
}
