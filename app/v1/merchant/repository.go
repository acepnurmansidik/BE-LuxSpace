package merchant

import "gorm.io/gorm"

type Repository interface {
	Save(merchant Merchant) (Merchant, error)
	FindByName(name string) (Merchant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(merchant Merchant) (Merchant, error) {
	err := r.db.Create(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repository) FindByName(name string) (Merchant, error) {
	var merchant Merchant
	err := r.db.Where("merchant_name = ?", name).Find(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}
