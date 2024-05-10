package service

import (
	"testing"
	"rub_buddy/features/pickup_transaction"
	"rub_buddy/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock untuk PickupTransactionDataInterface
type MockPickupTransactionData struct {
	mock.Mock
}

func (m *MockPickupTransactionData) CreatePickupTransaction(newData pickuptransaction.PickupTransaction) (*pickuptransaction.PickupTransactionCreate, error) {
	args := m.Called(newData)
	return args.Get(0).(*pickuptransaction.PickupTransactionCreate), args.Error(1)
}

func (m *MockPickupTransactionData) GetAllPickupTransaction(userId uint) ([]pickuptransaction.PickupTransactionInfo, error) {
	args := m.Called(userId)
	return args.Get(0).([]pickuptransaction.PickupTransactionInfo), args.Error(1)
}

func (m *MockPickupTransactionData) GetPickupTransactionByID(id int) (pickuptransaction.PickupTransactionInfo, error) {
	args := m.Called(id)
	return args.Get(0).(pickuptransaction.PickupTransactionInfo), args.Error(1)
}

func TestCreatePickupTransaction(t *testing.T) {
	mockData := new(MockPickupTransactionData)
	service := New(mockData)
	testData := pickuptransaction.PickupTransaction{PickupRequestID: 1, TpsID: 1}

	mockData.On("CreatePickupTransaction", testData).Return(&pickuptransaction.PickupTransactionCreate{}, nil)

	result, err := service.CreatePickupTransaction(testData)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	mockData.AssertExpectations(t)
}

func TestCreatePickupTransactionError(t *testing.T) {
	mockData := new(MockPickupTransactionData)
	service := New(mockData)
	testData := pickuptransaction.PickupTransaction{PickupRequestID: 0, TpsID: 0}

	result, err := service.CreatePickupTransaction(testData)
	assert.Nil(t, result)
	assert.Equal(t, constant.ErrPickupTransactionEmptyInput, err)
}

func TestGetAllPickupTransaction(t *testing.T) {
	mockData := new(MockPickupTransactionData)
	service := New(mockData)
	userId := uint(1)
	expectedResult := []pickuptransaction.PickupTransactionInfo{{}}

	mockData.On("GetAllPickupTransaction", userId).Return(expectedResult, nil)

	result, err := service.GetAllPickupTransaction(userId)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	mockData.AssertExpectations(t)
}

func TestGetPickupTransactionByID(t *testing.T) {
	mockData := new(MockPickupTransactionData)
	service := New(mockData)
	id := 1
	expectedResult := pickuptransaction.PickupTransactionInfo{}

	mockData.On("GetPickupTransactionByID", id).Return(expectedResult, nil)

	result, err := service.GetPickupTransactionByID(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	mockData.AssertExpectations(t)
}

