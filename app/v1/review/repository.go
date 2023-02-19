package review

import (
	"gorm.io/gorm"
)

type Repository interface {
	SaveUserReview(review Review) (Review, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveUserReview(review Review) (Review, error) {
	err := r.db.Create(&review).Error
	if err != nil {
		return review, err
	}

	return review, nil
}
