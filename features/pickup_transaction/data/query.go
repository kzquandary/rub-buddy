package data

import (
	"math/rand"
	"rub_buddy/features/chat"
	pickuptransaction "rub_buddy/features/pickup_transaction"
	"time"

	"gorm.io/gorm"
)

type PickupTransactionData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) pickuptransaction.PickupTransactionDataInterface {
	return &PickupTransactionData{
		DB: db,
	}
}

func (data *PickupTransactionData) CreatePickupTransaction(newPickupTransaction pickuptransaction.PickupTransaction) (*pickuptransaction.PickupTransactionCreate, error) {
	newPickupTransaction.PickupTime = time.Now()
	newPickupTransaction.CreatedAt = time.Now()
	newPickupTransaction.UpdatedAt = time.Now()
	newPickupTransaction.ID = uint(rand.Intn(900000) + 100000)
	
	chat := chat.Chat{
		PickupTransactionID: newPickupTransaction.ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ID:                  uint(rand.Intn(900000) + 100000),
	}

	if query := data.DB.Create(&newPickupTransaction); query.Error != nil {
		return nil, query.Error
	}

	if createChat := data.DB.Table("chats").Create(&chat); createChat.Error != nil {
		return nil, createChat.Error
	}

	if updateStatus := data.DB.Table("pickup_requests").Where("id = ?", newPickupTransaction.PickupRequestID).Update("status", "Accepted"); updateStatus.Error != nil {
		return nil, updateStatus.Error
	}

	var TransactionCreateInfo = pickuptransaction.PickupTransactionCreate{
		ID:              newPickupTransaction.ID,
		PickupRequestID: newPickupTransaction.PickupRequestID,
		CollectorID:     newPickupTransaction.CollectorID,
		TpsID:           newPickupTransaction.TpsID,
		PickupTime:      newPickupTransaction.PickupTime,
		ChatID:          chat.ID,
	}
	return &TransactionCreateInfo, nil
}

func (data *PickupTransactionData) GetAllPickupTransaction(userId uint) ([]pickuptransaction.PickupTransaction, error) {
	var pickupTransactions []pickuptransaction.PickupTransaction
	err := data.DB.Where("collector_id = ?", userId).Find(&pickupTransactions).Error
	return pickupTransactions, err
}

func (data *PickupTransactionData) GetPickupTransactionByID(id int) (pickuptransaction.PickupTransaction, error) {
	var pickupTransaction pickuptransaction.PickupTransaction
	err := data.DB.Where("id = ?", id).First(&pickupTransaction).Error
	return pickupTransaction, err
}
