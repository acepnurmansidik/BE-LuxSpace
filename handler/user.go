package handler

import (
	"LuxSpace/app/v1/user"
	"LuxSpace/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var inputData user.CreateUserInput
	// tangkap data
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var validEmail user.CheckEmailInput
	validEmail.Email = inputData.Email

	// lakukan validasi jika email sudah terdaftar
	isEmailRegister, err := h.service.IsEmailAvailable(validEmail)
	if err != nil {
		response := helper.APIResponse("Failed fetch data user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// cek jika emailnya ada maka kirim response email sudah terdaftar
	if isEmailRegister {
		response := helper.APIResponse("Email has been register", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// simpan ke databse jika sudha benar
	newUser, err := h.service.RegisterUser(inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Registration success", http.StatusOK, "success", user.FormatterUserLogin(newUser))
	c.JSON(http.StatusOK, response)
}
