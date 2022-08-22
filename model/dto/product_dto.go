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
