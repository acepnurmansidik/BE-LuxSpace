package handler

import (
	"LuxSpace/app/v1/category"
	"LuxSpace/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	service category.Service
}

func NewCategoryHandler(service category.Service) *categoryHandler {
	return &categoryHandler{service}
}

func (h *categoryHandler) GetCategorys(c *gin.Context) {
	queryName := c.Query("name")

	newCategory, err := h.service.GetCategories(queryName)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data category", http.StatusBadRequest, "error", nil)
		// kirim response
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Category", http.StatusOK, "success", category.FormatCategories(newCategory))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetDetailCategory(c *gin.Context) {
	var inputID category.CategoryDetailInput
	// ambil id nya
	err := c.ShouldBindUri(&inputID)

	// cek jika tidak ada id yang dikirim
	if err != nil {
		response := helper.APIResponse("Failed fecth data from parameter", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	newCategory, err := h.service.GetDetailCategory(inputID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail category", http.StatusOK, "success", category.FormatCategory(newCategory))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CreateCategoryInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("Failed create category", http.StatusBadRequest, "error", nil)
		// kirim response json nya
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCategory, err := h.service.CreateCategory(input)
	if err != nil {
		response := helper.APIResponse("Failed create category", http.StatusBadRequest, "error", nil)
		// kirim response json nya
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Category created", http.StatusOK, "success", category.FormatCategory(newCategory))
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var inputID category.CategoryDetailInput
	// tangkap id urinya
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("Failed to fetch data category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	//tangkap body yang akan di updatenya
	var inputData category.CreateCategoryInput
	err = c.ShouldBind(&inputData)

	if err != nil {
		response := helper.APIResponse("Failed to fetch data category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	// update data
	newCategory, err := h.service.UpdateCategory(inputID, inputData)

	if err != nil {
		response := helper.APIResponse("Failed get data category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	response := helper.APIResponse("Category update", http.StatusOK, "success", category.FormatCategory(newCategory))
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	var inputID category.CategoryDetailInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helper.APIResponse("Failed fetch data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCategory, err := h.service.DeleteCategory(inputID)
	if err != nil {
		response := helper.APIResponse("Data not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse("Category has been deleted", http.StatusOK, "error", category.FormatCategory(newCategory))
	c.JSON(http.StatusOK, response)
}
