package handler

import (
	"LuxSpace/app/v1/courir"
	"LuxSpace/helper"
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
