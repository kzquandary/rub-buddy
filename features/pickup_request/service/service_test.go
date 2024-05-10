package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"rub_buddy/constant"
	pickuprequest "rub_buddy/features/pickup_request"
)

// Mock data interface
type MockPickupRequestData struct {
	mock.Mock
}

func (m *MockPickupRequestData) CreatePickupRequest(newPickupRequest *pickuprequest.PickupRequest) (*pickuprequest.PickupRequest, error) {
	args := m.Called(newPickupRequest)
	return args.Get(0).(*pickuprequest.PickupRequest), args.Error(1)
}

func (m *MockPickupRequestData) GetAllPickupRequest() ([]pickuprequest.PickupRequestInfo, error) {
	args := m.Called()
	return args.Get(0).([]pickuprequest.PickupRequestInfo), args.Error(1)
}

func (m *MockPickupRequestData) GetPickupRequestByID(id int) (pickuprequest.PickupRequestInfo, error) {
	args := m.Called(id)
	return args.Get(0).(pickuprequest.PickupRequestInfo), args.Error(1)
}

func (m *MockPickupRequestData) DeletePickupRequestByID(id int, userID uint) error {
	args := m.Called(id, userID)
	return args.Error(0)
}

// Test cases
func TestCreatePickupRequest(t *testing.T) {
	mockData := new(MockPickupRequestData)
	service := New(mockData)
	testRequest := pickuprequest.PickupRequest{
		Weight:      10,
		Description: "Test Description",
		Image:       "test_image.png",
	}

	mockData.On("CreatePickupRequest", &testRequest).Return(&testRequest, nil)

	result, err := service.CreatePickupRequest(testRequest)
	assert.Nil(t, err)
	assert.Equal(t, &testRequest, result)
	mockData.AssertExpectations(t)
}

func TestCreatePickupRequestFailValidation(t *testing.T) {
	service := New(new(MockPickupRequestData))
	testRequest := pickuprequest.PickupRequest{}

	result, err := service.CreatePickupRequest(testRequest)
	assert.Nil(t, result)
	assert.Equal(t, constant.ErrPickupRequestEmptyInput, err)
}

func TestGetAllPickupRequest(t *testing.T) {
	mockData := new(MockPickupRequestData)
	service := New(mockData)
	expectedResult := []pickuprequest.PickupRequestInfo{
		{
			ID:          1,
			Weight:      15,
			Description: "Sample",
		},
	}

	mockData.On("GetAllPickupRequest").Return(expectedResult, nil)

	result, err := service.GetAllPickupRequest()
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	mockData.AssertExpectations(t)
}

func TestGetPickupRequestByID(t *testing.T) {
	mockData := new(MockPickupRequestData)
	service := New(mockData)
	expectedResult := pickuprequest.PickupRequestInfo{
		ID:          1,
		Weight:      15,
		Description: "Sample",
	}

	mockData.On("GetPickupRequestByID", 1).Return(expectedResult, nil)

	result, err := service.GetPickupRequestByID(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	mockData.AssertExpectations(t)
}

func TestDeletePickupRequestByID(t *testing.T) {
	mockData := new(MockPickupRequestData)
	service := New(mockData)

	// Setup mock expectations
	mockData.On("GetPickupRequestByID", 1).Return(pickuprequest.PickupRequestInfo{ID: 1}, nil)
	// Pastikan tipe data untuk userID adalah uint
	mockData.On("DeletePickupRequestByID", 1, uint(1)).Return(nil)

	// Call the method
	err := service.DeletePickupRequestByID(1, 1)

	// Assertions
	assert.Nil(t, err)
	mockData.AssertExpectations(t)
}

func TestDeletePickupRequestByIDNotFound(t *testing.T) {
	mockData := new(MockPickupRequestData)
	service := New(mockData)

	// Setup mock expectations for not found scenario
	mockData.On("GetPickupRequestByID", 1).Return(pickuprequest.PickupRequestInfo{ID: 0}, nil)

	// Call the method
	err := service.DeletePickupRequestByID(1, 1)

	// Assertions
	assert.Equal(t, constant.ErrPickupRequestNotFound, err)
	mockData.AssertExpectations(t)
}

func TestDeletePickupRequestByIDError(t *testing.T) {
	mockData := new(MockPickupRequestData)
	service := New(mockData)

	// Setup mock expectations for error scenario
	mockData.On("GetPickupRequestByID", 1).Return(pickuprequest.PickupRequestInfo{ID: 1}, nil)
	mockData.On("DeletePickupRequestByID", 1, uint(1)).Return(constant.ErrPickupRequestDelete)

	// Call the method
	err := service.DeletePickupRequestByID(1, 1)

	// Assertions
	assert.Equal(t, constant.ErrPickupRequestDelete, err)
	mockData.AssertExpectations(t)
}
