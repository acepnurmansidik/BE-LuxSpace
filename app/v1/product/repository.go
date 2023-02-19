package product

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	SaveProduct(product Product) (Product, error)
	SaveProductImage(productImage ProductImages) (ProductImages, error)
	FindAllByMerchant(merchantID int) ([]Product, error)
	Destroy(product Product) (Product, error)
	FindByID(ID int) (Product, error)
	Update(product Product) (Product, error)
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

func (r *repository) Destroy(product Product) (Product, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product
	err := r.db.Preload("Category").Preload("Merchant").Preload("ProductImages").Where("ID = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
