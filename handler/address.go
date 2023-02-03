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

	response := helper.APIResponse("Address list", http.StatusOK, "success", listAddress)
	c.JSON(http.StatusOK, response)
}
