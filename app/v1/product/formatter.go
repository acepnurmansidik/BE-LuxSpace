package product

import (
	"LuxSpace/app/v1/category"
	"LuxSpace/app/v1/merchant"
)

// product image list
type FormatProductImage struct {
	Name     string `json:"name" binding:"required"`
	IsActive string `json:"is_active" binding:"required"`
}

func FormatterProductImage(image ProductImages) FormatProductImage {
	formatter := FormatProductImage{}
	formatter.Name = image.Name
	formatter.IsActive = image.IsDelete

	return formatter
}

func FormatterFormatProductImagesList(images []ProductImages) []FormatProductImage {
	formatter := []FormatProductImage{}

	for _, image := range images {
		formatter = append(formatter, FormatterProductImage(image))
	}

	return formatter
}

// format product list
type FormatMerchantProducts struct {
	ID          int                         `json:"id" binding:"required"`
	Title       string                      `json:"title" binding:"required"`
	Description string                      `json:"description" binding:"required"`
	Stock       int                         `json:"stock" binding:"required"`
	Weight      int                         `json:"weight" binding:"required"`
	Price       int                         `json:"price" binding:"required"`
	Category    category.CategoryFormatter  `json:"category" binding:"required"`
	Merchant    merchant.FormatMerchantInfo `json:"merchant" binding:"required"`
	Images      []FormatProductImage        `json:"images" binding:"required"`
}

func FormatterMerchantProducts(products Product) FormatMerchantProducts {
	formatter := FormatMerchantProducts{}
	formatter.ID = products.ID
	formatter.Title = products.Title
	formatter.Description = products.Description
	formatter.Stock = products.Stock
	formatter.Weight = products.Weight
	formatter.Price = products.Price
	formatter.Category = category.FormatCategory(products.Category)
	formatter.Merchant = merchant.FormatterMerchantInfo(products.Merchant)
	formatter.Images = FormatterFormatProductImagesList(products.ProductImages)

	return formatter
}

func FormatterMerchantProductsList(products []Product) []FormatMerchantProducts {
	var productList []FormatMerchantProducts

	for _, prod := range products {
		productList = append(productList, FormatterMerchantProducts(prod))
	}

	return productList
}
