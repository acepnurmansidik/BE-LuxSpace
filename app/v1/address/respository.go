package address

import "gorm.io/gorm"

type Repository interface {
	FindAll(ID int) ([]Address, error)
	Save(address Address) (Address, error)
	Update(address Address) (Address, error)
	FindByID(addressID int) (Address, error)
	Destroy(address Address) (Address, error)
	MarkAllUserAddressNonPrimary(userID int) (bool, error)
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

func (r *repository) FindByID(addressID int) (Address, error) {
	var address Address
	err := r.db.Where("ID = ?", addressID).Find(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *repository) Destroy(address Address) (Address, error) {
	err := r.db.Delete(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *repository) MarkAllUserAddressNonPrimary(userID int) (bool, error) {
	err := r.db.Model(&Address{}).Where("user_id = ?", userID).Update("is_primary", "false").Error
	if err != nil {
		return false, err
	}

	return true, nil
}
