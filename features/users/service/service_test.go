package service

import (
	"fmt"
	"rub_buddy/constant"
	"rub_buddy/features/users"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockUserData adalah objek mock untuk UserDataInterface
type MockUserData struct {
	Users   map[string]*users.User
	Err     error
	Updated bool
}
type JWTInterface interface {
	GenerateJWT(userID uint, role, email, address string) (string, error)
	CheckID(token string) (uint, error)
}

type MockJWT struct {
	Token string
	Err   error
}

// ExtractToken implements helper.JWTInterface.
func (m *MockJWT) ExtractToken(token *jwt.Token) map[string]interface{} {
	panic("unimplemented")
}

// GenerateToken implements helper.JWTInterface.
func (m *MockJWT) GenerateToken(id uint, role string, email string, address string) string {
	panic("unimplemented")
}

// GetID implements helper.JWTInterface.
func (m *MockJWT) GetID(c echo.Context) (uint, error) {
	panic("unimplemented")
}

// ValidateToken implements helper.JWTInterface.
func (m *MockJWT) ValidateToken(token string) (*jwt.Token, error) {
	panic("unimplemented")
}

func (m *MockJWT) GenerateJWT(userID uint, role, email, address string) (string, error) {
	if m.Err != nil {
		return "", m.Err
	}
	return m.Token, nil
}

func (m *MockJWT) CheckID(ctx echo.Context) interface{} {
	// Implement this method to simulate checking ID from JWT
	// For testing purpose, you can return a static value
	// or check if the token is valid
	return 1 // Static value for testing
}

func NewMockUserData() *MockUserData {
	return &MockUserData{
		Users: make(map[string]*users.User), // Initialize Users map
	}
}

func (m *MockUserData) Login(email string, password string) (*users.User, error) {
	user, ok := m.Users[email]
	if !ok {
		return nil, constant.UserNotFound
	}
	if user.Password != password {
		return nil, constant.UserNotFound
	}
	return user, m.Err
}

func (m *MockUserData) Register(user users.User) (*users.User, error) {
	// Implementation to mock user registration
	// You can customize it as needed for your tests

	return &user, m.Err
}

func (m *MockUserData) UpdateUser(user *users.UserUpdate) error {
	// Implementation to mock user update
	// You can customize it as needed for your tests
	m.Users[user.Email] = &users.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Gender:   user.Gender,
	}
	m.Updated = true
	return m.Err
}

func (m *MockUserData) GetUserByEmail(email string) (*users.User, error) {
	user, ok := m.Users[email]
	if !ok {
		return nil, constant.UserNotFound
	}
	return user, m.Err
}

func TestLoginEmptyInput(t *testing.T) {
	service := New(&MockUserData{}, &MockJWT{})
	_, err := service.Login("", "")
	assert.Error(t, err)
	assert.Equal(t, constant.ErrLoginEmptyInput, err)
}

func TestLoginErrorJWT(t *testing.T) {
	mockUserData := &MockUserData{
		Users: map[string]*users.User{
			"john@example.com": {
				Email:    "john@example.com",
				Password: "password123",
			},
		},
	}
	mockJWT := &MockJWT{Err: fmt.Errorf("JWT error")}
	service := New(mockUserData, mockJWT)
	_, err := service.Login("john@example.com", "password123")
	assert.Error(t, err)
}

func TestRegisterUser(t *testing.T) {
	mockUserData := &MockUserData{Err: nil}
	service := New(mockUserData, &MockJWT{})
	newUser := users.User{
		Email:    "new@example.com",
		Password: "newpassword",
		Name:     "New User",
		Address:  "New Address",
		Gender:   "Male",
	}
	_, err := service.Register(newUser)
	assert.NoError(t, err)
}

func TestUpdateUserNotFound(t *testing.T) {
	mockUserData := NewMockUserData() 
	mockUserData.Err = constant.UserNotFound
	service := New(mockUserData, &MockJWT{})
	userUpdate := &users.UserUpdate{ID: 1}
	err := service.UpdateUser(userUpdate)
	assert.Error(t, err)
	assert.Equal(t, constant.UserNotFound, err)
}


func TestGetUserByEmailNotFound(t *testing.T) {
	mockUserData := &MockUserData{Err: constant.UserNotFound}
	service := New(mockUserData, &MockJWT{})
	_, err := service.GetUserByEmail("notfound@example.com")
	assert.Error(t, err)
	assert.Equal(t, constant.UserNotFound, err)
}
func TestLoginSuccess(t *testing.T) {
	mockUserData := &MockUserData{
		Users: map[string]*users.User{
			"john@example.com": {
				Email:    "john@example.com",
				Password: "password123",
			},
		},
	}
	mockJWT := &MockJWT{Token: "valid.token.here"}
	service := New(mockUserData, mockJWT)
	user, err := service.Login("john@example.com", "password123")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestRegisterUserExists(t *testing.T) {
	mockUserData := &MockUserData{
		Users: map[string]*users.User{
			"existing@example.com": {
				Email:    "existing@example.com",
				Password: "password123",
			},
		},
		Err: constant.ErrRegisterUserExists,
	}
	service := New(mockUserData, &MockJWT{})
	newUser := users.User{
		Email:    "existing@example.com",
		Password: "newpassword",
	}
	_, err := service.Register(newUser)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrRegisterEmptyInput, err)
}

func TestUpdateUserSuccess(t *testing.T) {
	mockUserData := NewMockUserData()
	mockUserData.Users["john@example.com"] = &users.User{
		ID:       1,
		Email:    "john@example.com",
		Password: "password123",
	}
	service := New(mockUserData, &MockJWT{})
	userUpdate := &users.UserUpdate{
		ID:       1,
		Email:    "john@example.com",
		Password: "newpassword123",
	}
	err := service.UpdateUser(userUpdate)
	assert.NoError(t, err)
	assert.True(t, mockUserData.Updated)
}

func TestGetUserByEmailSuccess(t *testing.T) {
	mockUserData := &MockUserData{
		Users: map[string]*users.User{
			"john@example.com": {
				ID:       1,
				Email:    "john@example.com",
				Password: "password123",
			},
		},
	}
	service := New(mockUserData, &MockJWT{})
	user, err := service.GetUserByEmail("john@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "john@example.com", user.Email)
}

func TestJWTErrorOnLogin(t *testing.T) {
	mockUserData := &MockUserData{
		Users: map[string]*users.User{
			"john@example.com": {
				Email:    "john@example.com",
				Password: "password123",
			},
		},
	}
	mockJWT := &MockJWT{Err: fmt.Errorf("JWT generation failed")}
	service := New(mockUserData, mockJWT)
	_, err := service.Login("john@example.com", "password123")
	assert.Error(t, err)
	assert.Equal(t, "JWT generation failed", err.Error())
}