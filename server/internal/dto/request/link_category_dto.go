package request

type LinkCategoryDTO struct {
	CategoryID int `json:"category_id" validate:"required,gt=0"`
}
