package product

type FormatProduct struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Weight      int    `json:"weight" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	CategoryId  int    `json:"category_id" binding:"required"`
	MerchantId  int    `json:"merchant_id" binding:"required"`
}

func FormatterProduct(product Product) FormatProduct {
	formatter := FormatProduct{}
	formatter.ID = product.ID
	formatter.Description = product.Description
	formatter.Title = product.Title
	formatter.Stock = product.Stock
	formatter.Weight = product.Weight
	formatter.Price = product.Price
	formatter.CategoryId = product.CategoryId
	formatter.MerchantId = product.MerchantId

	return formatter
}
