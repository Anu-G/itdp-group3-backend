package dto

type ProductRequest struct {
	AccountID   string `json:"account_id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Description string `json:"description"`

	DetailMediaProducts []DetailMediaProduct `json:"detail_media_products"`
}

type DetailMediaProduct struct {
	MediaLink string `json:"media_link"`
}
