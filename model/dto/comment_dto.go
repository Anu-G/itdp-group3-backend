package dto

type RequestCreateComment struct {
	FeedID      uint   `json:"feed_id,string"`
	AccountID   uint   `json:"account_id,string"`
	CommentFill string `json:"comment_fill"`
	DisplayName string `json:"display_name"`
	ProfileImage string `json:"profile_image"`
}

type RequestDeleteComment struct {
	CommentID uint `json:"id"`
}
