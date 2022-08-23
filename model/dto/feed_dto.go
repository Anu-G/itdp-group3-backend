package dto

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
