package helper

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func APIResponse(message string, code int, status string, data interface{}) *Response {
	// masukan meta nya
	meta := Meta{Message: message, Status: status, Code: code}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return &jsonResponse
}

func ImagesValidation(imageInput string) (bool, error) {
	// set extention image apa aja yang di perbolehkan
	imagesExt := []string{"jpg", "jpeg", "png"}
	newImage := strings.Split(imageInput, ".")
	for _, image := range imagesExt {
		if strings.ToLower(newImage[1]) == image {
			return true, nil
		}
	}

	return false, errors.New("image extention must PNG, JPG, JPEG")
}

func FormatValidationError(err error) []string {
	// var utk menampung error
	var errors []string

	// ubah error menjadi error validatior
	for _, e := range err.(validator.ValidationErrors) {
		// simpan setiap error string validator ke dalam slice errors
		errors = append(errors, e.Error())
	}

	return errors
}
