package handler

import (
	"LuxSpace/app/v1/address"
	"LuxSpace/app/v1/user"
	"LuxSpace/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addressHandler struct {
	service address.Service
}

func NewAddressHandler(service address.Service) *addressHandler {
	return &addressHandler{service}
}

func (h *addressHandler) GetAllAddress(c *gin.Context) {
	// ambil id user yang login
	valueUser, _ := c.Get("currentUser")
	userID := valueUser.(user.FormatUserHeader)

	var inputUserID address.AddressUserInput
	inputUserID.ID = userID.ID

	listAddress, err := h.service.GetAllAddress(inputUserID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Address list", http.StatusOK, "success", address.FormatterAddressList(listAddress))
	c.JSON(http.StatusOK, response)
}

func (h *addressHandler) CreateAddress(c *gin.Context) {
	// ambil semua data di body
	var inputData address.CreateAddressInput
	err := c.ShouldBindJSON(&inputData)

	if err != nil {
		response := helper.APIResponse("Failed create address", http.StatusBadRequest, "error", nil)
		// kirim response json nya
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil id user
	value, _ := c.Get("currentUser")
	userLogin := value.(user.FormatUserHeader)

	inputData.UserId = userLogin.ID

	// simpan data address ke database
	newAddress, err := h.service.CreateAddress(inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Address created", http.StatusOK, "success", address.FormatterAddress(newAddress))
	c.JSON(http.StatusOK, response)
	return
}
