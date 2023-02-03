package address

import "gorm.io/gorm"

type Repository interface {
	FindAll(ID int) ([]Address, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(ID int) ([]Address, error) {
	var address []Address
	err := r.db.Where("user_id = ?", ID).Find(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}
