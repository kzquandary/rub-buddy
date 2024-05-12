package data

import (
	"math/rand"
	"time"

	"gorm.io/gorm"

	"rub_buddy/constant"
	"rub_buddy/constant/tablename"
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
	newPickupRequest.Status = tablename.TablePickupRequestStatusPending
	newPickupRequest.CreatedAt = time.Now()
	newPickupRequest.UpdatedAt = time.Now()
	newPickupRequest.ID = uint(rand.Intn(900000) + 100000)
	err := data.DB.Create(newPickupRequest).Error
	if err != nil {
		return nil, constant.ErrPickupRequestCreate
	}
	return newPickupRequest, nil
}

func (data *PickupRequestData) GetAllPickupRequest(role string, id uint) ([]pickuprequest.PickupRequestInfo, error) {
	var pickupRequests []pickuprequest.PickupRequestInfo
	var err error

	if role == "Collector" {
		err = data.DB.Table(tablename.PickupRequestTableName).
			Select("pickup_requests.id, pickup_requests.user_id, users.name as user_name, pickup_requests.weight, pickup_requests.address, pickup_requests.description, pickup_requests.earnings, pickup_requests.image, pickup_requests.status").
			Joins("JOIN users ON pickup_requests.user_id = users.id").
			Where("pickup_requests.status = ?", tablename.TablePickupRequestStatusPending).
			Scan(&pickupRequests).Error
	} else if role == "User" {
		err = data.DB.Table(tablename.PickupRequestTableName).
			Select("pickup_requests.id, pickup_requests.user_id, users.name as user_name, pickup_requests.weight, pickup_requests.address, pickup_requests.description, pickup_requests.earnings, pickup_requests.image, pickup_requests.status").
			Joins("JOIN users ON pickup_requests.user_id = users.id").
			Where("pickup_requests.user_id = ?", id).
			Scan(&pickupRequests).Error
	} else {
		return nil, constant.ErrPickupRequestGet
	}

	if err != nil {
		return nil, constant.ErrPickupRequestGet
	}
	return pickupRequests, nil
}


func (data *PickupRequestData) GetPickupRequestByID(id int) (pickuprequest.PickupRequestInfo, error) {
	var pickupRequest pickuprequest.PickupRequestInfo
	err := data.DB.Table(tablename.PickupRequestTableName).
		Select("pickup_requests.id, pickup_requests.user_id, users.name as user_name, pickup_requests.weight, pickup_requests.address, pickup_requests.description, pickup_requests.earnings, pickup_requests.image, pickup_requests.status").
		Joins("JOIN users ON pickup_requests.user_id = users.id").
		Where("pickup_requests.id = ?", id).
		Scan(&pickupRequest).Error
	if err != nil {
		return pickuprequest.PickupRequestInfo{}, constant.ErrPickupRequestGet
	}
	return pickupRequest, nil
}

func (data *PickupRequestData) DeletePickupRequestByID(id int, userID uint) error {
	err := data.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&pickuprequest.PickupRequest{}).Error
	if err != nil {
		return constant.ErrPickupRequestDelete
	}
	return nil
}
