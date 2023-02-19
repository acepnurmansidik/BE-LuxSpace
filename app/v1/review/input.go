package review

type CreateReviewInput struct {
	Rate      int    `form:"rate" binding:"required"`
	Comments  string `form:"comments" binding:"required"`
	ProductID int    `form:"product_id" binding:"required"`
	UserId    int    `form:"user_id"`
}
