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

func (h *addressHandler) UpdateAddress(c *gin.Context) {
	// tagkap id table addresnya
	var inputID address.AddressDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// tangkap data yang akan di updatenya dari body
	var inputData address.CreateAddressInput
	err = c.ShouldBind(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil user id nya
	value, _ := c.Get("currentUser")
	userID := value.(user.FormatUserHeader)
	// mapping data user id ke inputan
	inputData.UserId = userID.ID

	// update data address
	newAddress, err := h.service.UpdateAddress(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Address has been updated", http.StatusOK, "success", address.FormatterAddress(newAddress))
	c.JSON(http.StatusOK, response)
}
