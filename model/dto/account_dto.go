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
