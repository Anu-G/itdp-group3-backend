package dto

import (
	"itdp-group3-backend/model/entity"
	"time"
)

type ReadPage struct {
	ID      uint `json:"account_id"`
	Cat     uint `json:"category"`
	Page    int  `json:"page"`
	PageLim int  `json:"page_lim"`
}

type DeleteFeed struct {
	ID uint `json:"feed_id"`
}

type DetailMediaFeed struct {
	MediaLink string `json:"media_link"`
}

type RequestFeed struct {
	AccountID   uint     `json:"account_ID"`
	CaptionPost string   `json:"caption_post"`
	MediaLinks  []string `json:"media_links"`
}

type ResponseFeed struct {
	AccountID   uint     `json:"account_ID"`
	CaptionPost string   `json:"caption_post"`
	MediaLinks  []string `json:"media_links"`
}

type RequestUpdateFeed struct {
	FeedID      uint     `json:"feed_ID"`
	CaptionPost string   `json:"caption_post"`
	MediaLinks  []string `json:"media_links"`
}

type FeedDetailRequest struct {
	AccountID        uint                   `json:"account_id"`
	PostID           uint                   `json:"post_id"`
	ProfileImage     string                 `json:"avatar"`
	CaptionPost      string                 `json:"caption_post"`
	CreatedAt        time.Time              `json:"created_at"`
	DetailMediaFeeds string                 `json:"detail_media_feed"`
	DisplayName      string                 `json:"display_name"`
	DetailComment    []entity.DetailComment `json:"detail_comment"`
	DetailLike       []entity.DetailLike    `json:"detail_like"`
}

type FeedDetailResponse struct {
	AccountID        uint                   `json:"account_id"`
	PostID           uint                   `json:"post_id"`
	ProfileImage     string                 `json:"avatar"`
	CaptionPost      string                 `json:"caption_post"`
	CreatedAt        time.Time              `json:"created_at"`
	DetailMediaFeeds []string               `json:"detail_media_feed"`
	DisplayName      string                 `json:"display_name"`
	DetailComment    []entity.DetailComment `json:"detail_comment"`
	DetailLike       []entity.DetailLike    `json:"detail_like"`
	TotalLike        int                    `json:"total_like"`
}

type LikeRequest struct {
	AccountID uint `json:"account_id"`
	FeedID    uint `json:"feed_id"`
}
