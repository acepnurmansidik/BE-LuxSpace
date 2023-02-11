package product

import (
	"LuxSpace/app/v1/category"
	"LuxSpace/app/v1/merchant"
)

type Product struct {
	ID            int
	Title         string
	Description   string
	Stock         int
	Weight        int
	Price         int
	CategoryId    int
	MerchantId    int
	ProductImages []ProductImages
	Category      category.Category
	Merchant      merchant.Merchant
}

type ProductImages struct {
	ID        int
	Name      string
	IsDelete  string
	ProductId int
}
