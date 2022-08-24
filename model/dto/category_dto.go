package dto

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type ReadCategoryResponse struct {
	CategoryNames []string `json:"category_names"`
}
