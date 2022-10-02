package dto

type ProductRequest struct {
	ProductID           string   `json:"product_id"`
	AccountID           string   `json:"account_id"`
	ProductName         string   `json:"product_name"`
	Price               string   `json:"price"`
	Description         string   `json:"description"`
	DetailMediaProducts []string `json:"detail_media_products"`
}

type ProductResponse struct {
	ProductID           string   `json:"product_id"`
	AccountID           string   `json:"account_id"`
	ProductName         string   `json:"product_name"`
	Price               string   `json:"price"`
	Description         string   `json:"description"`
	DetailMediaProducts []string `json:"detail_media_products"`
}

type SearchProductRequest struct {
	Keyword string `json:"keyword"`
}

type SearchProductResponse struct {
	ProductID           uint    `json:"product_id"`
	AccountID           string  `json:"account_id"`
	AccountAvatar       string  `json:"account_avatar"`
	AccountDisplayName  string  `json:"account_display_name"`
	ProductName         string  `json:"product_name"`
	ProductPrice        float64 `json:"product_price"`
	ProductDescription  string  `json:"product_description"`
	DetailMediaProducts string  `json:"detail_media_products"`
	CategoryID          string  `json:"category_id"`
}

type ProductDetailResponse struct {
	ProductID           uint     `json:"product_id"`
	AccountID           string   `json:"account_id"`
	ProfileImage        string   `json:"avatar"`
	Name                string   `json:"profile_name"`
	ProductName         string   `json:"product_name"`
	ProductPrice        float64  `json:"price"`
	Caption             string   `json:"caption"`
	DetailMediaProducts []string `json:"detail_media_products"`
	CategoryID          string  `json:"category_id"`
}
