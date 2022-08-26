package dto

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type ReadCategoryResponse struct {
	CategoryID    uint   `json:"category_id"`
	CategoryNames string `json:"category_names"`
}
