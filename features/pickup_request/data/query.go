package data

import (
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"rub_buddy/constant"
	pickuprequest "rub_buddy/features/pickup_request"
)

type PickupRequestData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) pickuprequest.PickupRequestDataInterface {
	return &PickupRequestData{
		DB: db,
	}
}

func (data *PickupRequestData) CreatePickupRequest(newPickupRequest *pickuprequest.PickupRequest) (*pickuprequest.PickupRequest, error) {
	newPickupRequest.Status = "Pending"
	newPickupRequest.CreatedAt = time.Now()
	newPickupRequest.UpdatedAt = time.Now()
	newPickupRequest.ID = uint(rand.Intn(900000) + 100000)
	return newPickupRequest, data.DB.Create(newPickupRequest).Error
}

func (data *PickupRequestData) GetAllPickupRequest() ([]pickuprequest.PickupRequest, error) {
	var pickupRequest []pickuprequest.PickupRequest
	err := data.DB.Where("status = ?", "Pending").Find(&pickupRequest).Error
	return pickupRequest, err
}

func (data *PickupRequestData) GetPickupRequestByID(id int) (pickuprequest.PickupRequest, error) {
	var pickupRequest pickuprequest.PickupRequest
	query := data.DB.Where("id = ?", id).Find(&pickupRequest)
	if query.RowsAffected == 0 {
		return pickupRequest, errors.New(constant.NotFound)
	}
	return pickupRequest, nil
}

func (data *PickupRequestData) DeletePickupRequestByID(id int, userID uint) error {
	return data.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&pickuprequest.PickupRequest{}).Error
}
