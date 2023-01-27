package category

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Category, error)
	FindByName(name string) ([]Category, error)
	FindByID(ID int) (Category, error)
	Save(category Category) (Category, error)
	Update(category Category) (Category, error)
	Destroy(category Category) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Category, error) {
	var category []Category

	err := r.db.Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindByID(ID int) (Category, error) {
	var category Category
	err := r.db.Where("ID = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindByName(name string) ([]Category, error) {
	var category []Category

	name = "%" + name + "%"
	err := r.db.Where("Name Like ?", name).Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Destroy(category Category) (Category, error) {
	err := r.db.Delete(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
