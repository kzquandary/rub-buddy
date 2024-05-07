package handler

type UserInfo struct {
	ID   uint   `json:"user_id"`
	Name string `json:"user_name"`
}
type PickupRequestInfo struct {
	ID   uint     `json:"pickup_request_id"`
	User UserInfo `json:"user"`

	Weight      float64 `json:"weight"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
	Earnings    float64 `json:"earnings"`
	Image       string  `json:"image"`
}
