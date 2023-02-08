package product

import "gorm.io/gorm"

type Repository interface {
	SaveProduct(product Product) (Product, error)
	SaveProductImage(productImage ProductImages) (ProductImages, error)
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
