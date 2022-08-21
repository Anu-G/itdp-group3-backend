package dto

type ReadPage struct {
	ID      uint `json:"account_id"`
	Page    int  `json:"page"`
	PageLim int  `json:"page_lim"`
}

type DeleteFeed struct {
	ID uint `json:"feed_id"`
}
