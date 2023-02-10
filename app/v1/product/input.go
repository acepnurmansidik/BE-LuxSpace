package product

type CreateProductInput struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Stock       int    `form:"stock" binding:"required"`
	Weight      int    `form:"weight" binding:"required"`
	Price       int    `form:"price" binding:"required"`
	CategoryID  int    `form:"category_id" binding:"required"`
	MerchantID  int
}

type CreateProductImagesInput struct {
	Name      string `form:"name" binding:"required"`
	IsPrimary string `form:"is_primary" binding:"required"`
	ProductId int    `form:"product_id" binding:"required"`
}

type ProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
