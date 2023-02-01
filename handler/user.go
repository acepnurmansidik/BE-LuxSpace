package handler

import (
	"LuxSpace/app/v1/user"
	"LuxSpace/auth"
	"LuxSpace/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	isEmailRegister, err := h.userService.IsEmailAvailable(validEmail)
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

	// simpan ke database jika sudha benar
	newUser, err := h.userService.RegisterUser(inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Registration success", http.StatusOK, "success", user.FormatterUserRegister(newUser))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var inputData user.LoginInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// cek user login
	loginUser, err := h.userService.Login(inputData)
	if err != nil {
		// cek jika ada error
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// generate token
	token, err := h.authService.GenerateToken(loginUser)
	if err != nil {
		// cek jika ada error
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get data user", http.StatusOK, "success", user.FormatterUserLogin(loginUser, token))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ActivateUser(c *gin.Context) {
	// ambil kode otp
	var codeOtp user.ActivateOtpInput
	err := c.ShouldBindUri(&codeOtp)
	if err != nil {
		response := helper.APIResponse("Failed fetch code", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// ambil emailnya
	var inputData user.CheckEmailInput
	err = c.ShouldBind(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch code", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.IsActivateUser(codeOtp, inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch code", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Account has been active", http.StatusOK, "success", user.FormatterUserRegister(newUser))
	c.JSON(http.StatusOK, response)
}
