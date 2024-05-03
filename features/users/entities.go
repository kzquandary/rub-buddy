package users

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Address   string        `json:"address"`
	Gender    string        `json:"gender"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *sql.NullTime `json:"deleted_at"`
}

type UserCredentials struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token any    `json:"token"`
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
	GetUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
}

type UserDataInterface interface {
	Register(user User) (*User, error)
	Login(email string, password string) (*User, error)
	GetUser(user *User) error
	UpdateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}
