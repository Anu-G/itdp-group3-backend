package dto

type RequestCreateComment struct {
	FeedID      uint   `json:"feed_id"`
	CommentFill string `json:"comment_fill"`
}

type RequestDeleteComment struct {
	CommentID uint `json:"id"`
}
