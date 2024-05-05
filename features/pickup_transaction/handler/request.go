package handler

type PickupTransactionInput struct {
	PickupRequestID uint `json:"pickup_request_id"`
	TpsID          uint `json:"tps_id"`
}