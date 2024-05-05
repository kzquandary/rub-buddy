package handler

type PickupTransactionInfo struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	CollectorID string  `json:"collector_id"`
	PickupTime  string  `json:"pickup_time"`
	TpsID       uint    `json:"tps_id"`
}