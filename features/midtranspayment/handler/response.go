package handler

type MidtransResponse struct {
	ID      string `json:"order_id"`
	UserID  uint   `json:"user_id"`
	Amount  int64  `json:"amount"`
	SnapURL string `json:"snap_url"`
}
