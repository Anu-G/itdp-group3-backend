package dto

type RegisterUserRequest struct {
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUserRegister struct {
}

type UpdateUserRequest struct {
	AccountID   uint   `json:"account_id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
