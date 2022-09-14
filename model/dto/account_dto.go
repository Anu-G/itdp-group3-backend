package dto

type FollowRequest struct {
	FollowerAccounID  uint `json:"follower_account_id"`
	FollowedAccountID uint `json:"followed_account_id"`
}

type UnfollowRequest struct {
	FollowRequest
}

type FollowListRequest struct {
	AccountID    uint `json:"account_id"`
	FollowStatus bool `json:"follow_status"`
}

type FollowListResponse struct {
	AccountID uint   `json:"account_id"`
	Username  string `json:"user_name"`
}

type ActivateBusinessAccountRequest struct {
	AccountID uint `json:"account_id"`
}

type AccountFillRequest struct {
	AccountID uint `json:"account_id"`
}

type GetAccountRequest struct {
	AccountID string `json:"account_id"`
}

type GetAccountResponse struct {
	Email string `json:"email"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}