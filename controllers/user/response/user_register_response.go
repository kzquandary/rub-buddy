package response

import "rub_buddy/entities"

type UserRegisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromUseCaseToRegister(user *entities.User) *UserRegisterResponse {
	return &UserRegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
