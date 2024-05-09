package handler

type UserInfo struct {
	ID   uint   `json:"user_id"`
	Name string `json:"user_name"`
}

type CollectorInfo struct {
	ID   uint   `json:"collector_id"`
	Name string `json:"collector_name"`
}

type ChatInfo struct {
	ID                  uint          `json:"chat_id"`
	PickupTransactionID uint          `json:"pickup_transaction_id"`
	User                UserInfo      `json:"user"`
	Collector           CollectorInfo `json:"collector"`
}
