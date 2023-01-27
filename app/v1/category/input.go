package category

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type CategoryDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
