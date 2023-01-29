package courir

type CourirDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCourirInput struct {
	Name string `json:"name" binding:"required"`
}
