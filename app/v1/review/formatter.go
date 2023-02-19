package review

import "LuxSpace/app/v1/user"

type FormateImageReview struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"image_name" binding:"required"`
}

func FormatterImageReview(inputData ReviewImages) FormateImageReview {
	formatter := FormateImageReview{}
	formatter.ID = inputData.ID
	formatter.Name = inputData.Name

	return formatter
}

func FormatterImagesListReview(imagesReview []ReviewImages) []FormateImageReview {
	formatter := []FormateImageReview{}

	for _, image := range imagesReview {
		formatter = append(formatter, FormatterImageReview(image))
	}

	return formatter
}

type FormatUserReview struct {
	ID     int    `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}

func FormatterUserReview(inputUser user.User) FormatUserReview {
	formatter := FormatUserReview{}

	formatter.ID = inputUser.ID
	formatter.Name = inputUser.Username
	formatter.Avatar = inputUser.Avatar

	return formatter
}

type FormatReview struct {
	ID       int                  `json:"review_id" binding:"required"`
	Rate     int                  `json:"rate" binding:"required"`
	Comments string               `json:"comments" binding:"required"`
	User     FormatUserReview     `json:"user_review" binding:"required"`
	Images   []FormateImageReview `json:"images_review" binding:"required"`
}

func FormatterReview(inputData Review) FormatReview {
	formatter := FormatReview{}
	formatter.ID = inputData.ID
	formatter.Rate = inputData.Rate
	formatter.Comments = inputData.Comments
	formatter.User = FormatterUserReview(inputData.User)
	formatter.Images = FormatterImagesListReview(inputData.Images)

	return formatter
}

func FormatterListReview(listReview []Review) []FormatReview {
	formatter := []FormatReview{}

	for _, reviewDetail := range listReview {
		formatter = append(formatter, FormatterReview(reviewDetail))
	}

	return formatter
}
