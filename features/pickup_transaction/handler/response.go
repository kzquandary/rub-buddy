package handler

type PickupTransactionCreate struct {
	ID              uint   `json:"id"`
	PickupRequestID uint   `json:"pickup_request_id"`
	PickupTime      string `json:"pickup_time"`
}
type PickupTransactionInfo struct {
	ID         uint          `json:"pickup_transaction_id"`
	User       UserInfo      `json:"user"`
	Collector  CollectorInfo `json:"collector"`
	PickupTime string        `json:"pickup_time"`
	Tps        TPSInfo       `json:"tps"`
}

type UserInfo struct {
	ID      uint   `json:"user_id"`
	Name    string `json:"user_name"`
	Address string `json:"user_address"`
}

type CollectorInfo struct {
	ID   uint   `json:"collector_id"`
	Name string `json:"collector_name"`
}

type TPSInfo struct {
	ID      uint   `json:"tps_id"`
	Name    string `json:"tps_name"`
	Address string `json:"tps_address"`
}
