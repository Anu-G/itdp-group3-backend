package dto

type RegisterUserRequest struct {
	Username string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
