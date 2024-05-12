package users

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Address   string
	Gender    string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *sql.NullTime
}

type UserUpdate struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Address   string
	Gender    string
	UpdatedAt time.Time
}

type UserCredentials struct {
	ID    uint
	Email string
	Token any
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newUser User) (*User, error)
	Login(email string, password string) (*UserCredentials, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *UserUpdate) error
}

type UserDataInterface interface {
	Register(user User) (*User, error)
	Login(email string, password string) (*User, error)
	UpdateUser(user *UserUpdate) error
	GetUserByEmail(email string) (*User, error)
}
