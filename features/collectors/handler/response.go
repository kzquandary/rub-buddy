package handler

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token any    `json:"token"`
}


type CollectorInfoResponse struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}


