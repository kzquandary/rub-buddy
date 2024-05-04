package collectors

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type Collectors struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Gender    string        `json:"gender"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *sql.NullTime `json:"deleted_at"`
}

type CollectorCredentials struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token any    `json:"token"`
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
	GetCollector(user *Collectors) error
	GetCollectorByEmail(email string) (*Collectors, error)
	UpdateCollector(user *Collectors) error
}

type CollectorDataInterface interface {
	Register(collector Collectors) (*Collectors, error)
	Login(email string, password string) (*Collectors, error)
	GetCollector(collector *Collectors) error
	UpdateCollector(collector *Collectors) error
	GetCollectorByEmail(email string) (*Collectors, error)
}
