package service

import (
	"errors"
	"rub_buddy/constant"
	"rub_buddy/features/collectors"
	"testing"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockCollectorData struct {
	Collectors   map[string]*collectors.Collectors
	Err          error
	Updated      bool
	EmailExists  bool
	UserNotFound bool
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

func NewMockCollectorData() *MockCollectorData {
	return &MockCollectorData{
		Collectors: make(map[string]*collectors.Collectors),
	}
}

func (m *MockCollectorData) Register(newCollector collectors.Collectors) (*collectors.Collectors, error) {
	if m.EmailExists {
		return nil, constant.ErrCollectorUserEmailExists
	}
	return &newCollector, m.Err
}

func (m *MockCollectorData) Login(email string, password string) (*collectors.Collectors, error) {
	if m.UserNotFound {
		return nil, constant.ErrCollectorUserNotFound
	}
	if m.Err != nil {
		return nil, m.Err
	}
	collector, ok := m.Collectors[email]
	if !ok {
		return nil, errors.New("collector not found")
	}
	return collector, nil
}

func (m *MockCollectorData) UpdateCollector(collector *collectors.CollectorUpdate) error {
	if m.UserNotFound {
		return constant.ErrCollectorUserNotFound
	}
	if m.Err != nil {
		return m.Err
	}
	m.Updated = true
	return nil
}

func (m *MockCollectorData) GetCollectorByEmail(email string) (*collectors.Collectors, error) {
	if m.UserNotFound {
		return nil, constant.ErrCollectorUserNotFound
	}
	collector, ok := m.Collectors[email]
	if !ok {
		return nil, errors.New("collector not found")
	}
	return collector, nil
}

func TestRegisterCollectorEmptyInput(t *testing.T) {
	service := New(NewMockCollectorData(), nil)
	_, err := service.Register(collectors.Collectors{})
	assert.Error(t, err)
	assert.Equal(t, constant.ErrRegisterEmptyInput, err)
}

func TestRegisterCollectorEmailExists(t *testing.T) {
	mockData := NewMockCollectorData()
	mockData.EmailExists = true
	service := New(mockData, nil)
	newCollector := collectors.Collectors{
		Email:    "existing@example.com",
		Password: "password123",
		Name:     "Existing Collector",
		Gender:   "Male",
	}
	_, err := service.Register(newCollector)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrCollectorUserEmailExists, err)
}

func TestLoginCollectorEmptyInput(t *testing.T) {
	service := New(NewMockCollectorData(), nil)
	_, err := service.Login("", "")
	assert.Error(t, err)
	assert.Equal(t, constant.ErrLoginEmptyInput, err)
}

func TestLoginCollectorNotFound(t *testing.T) {
	mockData := NewMockCollectorData()
	mockData.UserNotFound = true
	service := New(mockData, nil)
	_, err := service.Login("notfound@example.com", "password123")
	assert.Error(t, err)
	assert.Equal(t, constant.ErrCollectorUserNotFound, err)
}

func TestUpdateCollectorNotFound(t *testing.T) {
	mockData := NewMockCollectorData()
	mockData.UserNotFound = true
	service := New(mockData, nil)
	collectorUpdate := &collectors.CollectorUpdate{ID: 1}
	err := service.UpdateCollector(collectorUpdate)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrCollectorUserNotFound, err)
}

func TestGetCollectorByEmailNotFound(t *testing.T) {
	mockData := NewMockCollectorData()
	mockData.UserNotFound = true
	service := New(mockData, nil)
	_, err := service.GetCollectorByEmail("notfound@example.com")
	assert.Error(t, err)
	assert.Equal(t, constant.ErrCollectorUserNotFound, err)
}
func TestRegisterCollectorSuccess(t *testing.T) {
	mockData := NewMockCollectorData()
	service := New(mockData, nil)
	newCollector := collectors.Collectors{
		Email:    "new@example.com",
		Password: "newpassword",
		Name:     "New Collector",
		Gender:   "Male",
	}
	result, err := service.Register(newCollector)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLoginCollectorSuccess(t *testing.T) {
	mockData := NewMockCollectorData()
	mockData.Collectors["john@example.com"] = &collectors.Collectors{
		Email:    "john@example.com",
		Password: "password123",
		Name:     "John Doe",
		Gender:   "Male",
	}
	mockJWT := &MockJWT{Token: "mockToken"}
	service := New(mockData, mockJWT)
	result, err := service.Login("john@example.com", "password123")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "john@example.com", result.Email)
	assert.Equal(t, "mockToken", result.Token)
	assert.NotNil(t, result.ID)
}
