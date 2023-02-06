package merchant

import "gorm.io/gorm"

type Repository interface {
	Save(merchant Merchant) (Merchant, error)
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
