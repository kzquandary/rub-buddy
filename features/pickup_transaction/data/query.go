package data

import (
	"math/rand"
	"rub_buddy/constant"
	"rub_buddy/constant/tablename"
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
		return nil, constant.ErrPickupTransactionCreate
	}

	if createChat := data.DB.Table(tablename.ChatTableName).Create(&chat); createChat.Error != nil {
		return nil, constant.ErrPickupTransactionCreateChat
	}

	if updateStatus := data.DB.Table(tablename.PickupRequestTableName).Where("id = ?", newPickupTransaction.PickupRequestID).Update("status", tablename.TablePickupRequestStatusAccepted); updateStatus.Error != nil {
		return nil, constant.ErrPickupTransactionUpdateStatus
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

func (data *PickupTransactionData) GetAllPickupTransaction(userId uint) ([]pickuptransaction.PickupTransactionInfo, error) {
	var pickupTransactions []pickuptransaction.PickupTransactionInfo

	err := data.DB.Table(tablename.PickupTransactionTableName).
		Select("pickup_transactions.id, pickup_requests.user_id, users.name as user_name, users.address as user_address, collectors.name as collector_name, pickup_transactions.collector_id, pickup_transactions.pickup_time, pickup_transactions.tps_id, tps.name as tps_name, tps.address as tps_address").
		Joins("JOIN pickup_requests ON pickup_transactions.pickup_request_id = pickup_requests.id").
		Joins("JOIN users ON pickup_requests.user_id = users.id").
		Joins("JOIN collectors ON pickup_transactions.collector_id = collectors.id").
		Joins("JOIN tps ON pickup_transactions.tps_id = tps.id").
		Where("pickup_transactions.collector_id = ?", userId).
		Find(&pickupTransactions).Error

	if err != nil {
		return nil, constant.ErrPickupTransactionGet
	}

	return pickupTransactions, nil
}

func (data *PickupTransactionData) GetPickupTransactionByID(id int) (pickuptransaction.PickupTransactionInfo, error) {
	var pickupTransactions pickuptransaction.PickupTransactionInfo

	err := data.DB.Table(tablename.PickupTransactionTableName).
		Select("pickup_transactions.id, pickup_requests.user_id, users.name as user_name, users.address as user_address, collectors.name as collector_name, pickup_transactions.collector_id, pickup_transactions.pickup_time, pickup_transactions.tps_id, tps.name as tps_name, tps.address as tps_address").
		Joins("JOIN pickup_requests ON pickup_transactions.pickup_request_id = pickup_requests.id").
		Joins("JOIN users ON pickup_requests.user_id = users.id").
		Joins("JOIN collectors ON pickup_transactions.collector_id = collectors.id").
		Joins("JOIN tps ON pickup_transactions.tps_id = tps.id").
		Where("pickup_transactions.id = ?", id).
		Find(&pickupTransactions).Error

	if err != nil {
		return pickuptransaction.PickupTransactionInfo{}, constant.ErrPickupTransactionGet
	}

	return pickupTransactions, nil
}
