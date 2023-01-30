package handler

import (
	"LuxSpace/app/v1/courir"
	"LuxSpace/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type courirHandler struct {
	service courir.Service
}

func NewCourirHandler(service courir.Service) *courirHandler {
	return &courirHandler{service}
}

func (h *courirHandler) GetCourirs(c *gin.Context) {
	courirs, err := h.service.GetCourirs()
	if err != nil {
		response := helper.APIResponse("Failed to fetch data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get list courirs", http.StatusOK, "success", courir.FormatCourirs(courirs))
	c.JSON(http.StatusOK, response)
}

func (h *courirHandler) GetDetailCourir(c *gin.Context) {
	var inputID courir.CourirDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	courir, err := h.service.GetDetailCourir(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get detail data courir", http.StatusOK, "success", courir)
	c.JSON(http.StatusOK, response)
}

func (h *courirHandler) CreateCourir(c *gin.Context) {
	var input courir.CreateCourirInput
	// passing data dari inputan user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed create data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCourir, err := h.service.CreateCourir(input)
	if err != nil {
		response := helper.APIResponse("Failed create data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Courir has been created", http.StatusOK, "success", courir.FormatCourir(newCourir))
	c.JSON(http.StatusOK, response)
}

func (h *courirHandler) UpdateCourir(c *gin.Context) {
	var inputID courir.CourirDetailInput
	// ambil id kurirnya
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("disini 1")
		return
	}

	var inputData courir.CreateCourirInput
	// ambil data kurir yang akan di update/body nya
	err = c.ShouldBind(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		fmt.Println("disini 2")
		return
	}

	// simpan datanya
	newCourir, err := h.service.UpdateCourir(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed updated data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Courur updated", http.StatusOK, "success", courir.FormatCourir(newCourir))
	c.JSON(http.StatusOK, response)
}

func (h *courirHandler) DeleteCourir(c *gin.Context) {
	var inputID courir.CourirDetailInput
	// ambil id kurirnya
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed updated data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCourir, err := h.service.DeleteCourir(inputID)
	if err != nil {
		response := helper.APIResponse("Failed updated data courir", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Courir has been deleted", http.StatusOK, "success", courir.FormatCourir(newCourir))
	c.JSON(http.StatusOK, response)
}
