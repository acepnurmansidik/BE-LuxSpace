package handler

import (
	"LuxSpace/app/v1/merchant"
	"LuxSpace/app/v1/user"
	"LuxSpace/helper"
	"crypto/rand"
	"fmt"
	"math/big"
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
		if err.Error() != "" {
			response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("Failed fetch data merchant", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Merhact created", http.StatusOK, "success", merchant.FormatterMerchant(newMerchant))
	c.JSON(http.StatusOK, response)
}

func (h *merchantHandler) UploadImageMerchant(c *gin.Context) {
	// ambil file imagenya
	file, err := c.FormFile("image_url")
	if err != nil {
		response := helper.APIResponse("Failed fetch file image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	imageValid, err := helper.ImagesValidation(file.Filename)
	if err != nil || !imageValid {
		response := helper.APIResponse("Please insert image extention(png, jpeg, jpg)", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusForbidden, response)
		return
	}

	// ambil data user yang ingin membuka merchant
	userLogin := c.MustGet("currentUser").(user.FormatUserHeader)

	// buat angka acak untuk me-rename name file
	randomNumb, _ := rand.Int(rand.Reader, big.NewInt(99999999999))
	// sambung nama file dengan angka acak tersebut, lalu set path lokasinya
	path := fmt.Sprintf("images/merchants/%v-%d-%s", userLogin.ID, randomNumb, file.Filename)

	// upload fie image
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed upload file image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newMerchant, err := h.service.UpdateMerchant(path, userLogin.ID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data merchant", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Merhact created", http.StatusOK, "success", merchant.FormatterUploadImageMerchant(newMerchant))
	c.JSON(http.StatusOK, response)
}
