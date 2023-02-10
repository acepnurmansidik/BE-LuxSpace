package product

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	SaveProduct(product Product) (Product, error)
	SaveProductImage(productImage ProductImages) (ProductImages, error)
	FindAllByMerchant(merchantID int) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveProduct(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) SaveProductImage(productImage ProductImages) (ProductImages, error) {
	err := r.db.Create(&productImage).Error
	if err != nil {
		return productImage, err
	}

	return productImage, nil
}

func (r *repository) FindAllByMerchant(merchantID int) ([]Product, error) {
	var products []Product
	err := r.db.Preload(clause.Associations).Where("merchant_id = ?", merchantID).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}
