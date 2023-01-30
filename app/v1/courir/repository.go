package courir

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Courir, error)
	FindByID(inputID int) (Courir, error)
	Save(courir Courir) (Courir, error)
	Update(courir Courir) (Courir, error)
	Destroy(courir Courir) (Courir, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Courir, error) {
	var courir []Courir

	err := r.db.Find(&courir).Error

	if err != nil {
		return courir, err
	}

	return courir, nil
}

func (r *repository) FindByID(inputID int) (Courir, error) {
	var courir Courir
	err := r.db.Where("ID = ?", inputID).Find(&courir).Error

	if err != nil {
		return courir, err
	}

	return courir, nil
}

func (r *repository) Save(courir Courir) (Courir, error) {
	err := r.db.Create(&courir).Error
	if err != nil {
		return courir, err
	}

	return courir, nil
}

func (r *repository) Update(courir Courir) (Courir, error) {
	err := r.db.Save(&courir).Error
	if err != nil {
		return courir, err
	}

	return courir, nil
}

func (r *repository) Destroy(courir Courir) (Courir, error) {
	err := r.db.Delete(&courir).Error
	if err != nil {
		return courir, err
	}

	return courir, nil
}
