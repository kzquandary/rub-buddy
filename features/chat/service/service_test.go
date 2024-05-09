package service

import (
	"rub_buddy/features/chat"
	"testing"

	"github.com/stretchr/testify/assert" 
)

type MockChatData struct {
	Chats []chat.ChatInfo
	Err   error
}

func (m *MockChatData) GetChat(id uint, role string) ([]chat.ChatInfo, error) {
	return m.Chats, m.Err
}

func TestGetChat(t *testing.T) {
	mockData := &MockChatData{
		Chats: []chat.ChatInfo{
			{ID: 1, PickupTransactionID: 100, UserID: 10, UserName: "User1", CollectorID: 20, CollectorName: "Collector1"},
			{ID: 2, PickupTransactionID: 101, UserID: 11, UserName: "User2", CollectorID: 21, CollectorName: "Collector2"},
		},
		Err: nil,
	}

	service := New(mockData)

	chats, err := service.GetChat(10, "User")

	assert.NoError(t, err)
	assert.Len(t, chats, 2)
	assert.Equal(t, chats[0].ID, uint(1))
	assert.Equal(t, chats[0].UserName, "User1")
	assert.Equal(t, chats[1].ID, uint(2))
	assert.Equal(t, chats[1].UserName, "User2")
}
