package address

import "gorm.io/gorm"

type Repository interface {
	FindAll(ID int) ([]Address, error)
	Save(address Address) (Address, error)
	Update(address Address) (Address, error)
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

func (r *repository) Save(address Address) (Address, error) {
	err := r.db.Create(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *repository) Update(address Address) (Address, error) {
	err := r.db.Save(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}
