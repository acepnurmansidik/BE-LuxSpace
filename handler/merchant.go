package handler

import (
	"LuxSpace/app/v1/merchant"
	"LuxSpace/app/v1/user"
	"LuxSpace/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type merchantHandler struct {
	service merchant.Service
}

func NewMerchantHandler(service merchant.Service) *merchantHandler {
	return &merchantHandler{service}
}

func (h *merchantHandler) CreateUserMerchant(c *gin.Context) {
	var inputData merchant.CreateMerchantInput
	// ambil datanya dari form
	err := c.Bind(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data merchant", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil data user yang ingin membuka merchant
	userLogin := c.MustGet("currentUser").(user.FormatUserHeader)

	// mapping id user yang login ke merchant
	inputData.UserId = userLogin.ID

	// lalu simpan ke database
	newMerchant, err := h.service.CreateMerchant(inputData)
	if err != nil {
		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed fetch data merchant", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Merhact created", http.StatusOK, "success", merchant.FormatterMerchant(newMerchant))
	c.JSON(http.StatusOK, response)
}
