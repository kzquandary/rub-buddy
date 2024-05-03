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

type UserInfoResponse struct {
	Data any `json:"data"`
}
