package handler

type MidtransRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}
