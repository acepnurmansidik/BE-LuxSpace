package handler

import (
	"LuxSpace/app/v1/product"
	"LuxSpace/helper"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) CreateProductMerchant(c *gin.Context) {
	var inputDataProduct product.CreateProductInput
	err := c.ShouldBind(&inputDataProduct)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// simpan ke database
	newProduct, err := h.service.CreateProduct(inputDataProduct)
	if err != nil {
		response := helper.APIResponse("Failed create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil file gambar dari form dengan bentuk array
	form, _ := c.MultipartForm()
	files := form.File["imagesUrl"]
	for _, file := range files {
		imageValid, err := helper.ImagesValidation(file.Filename)
		if err != nil || !imageValid {
			response := helper.APIResponse("Please insert image extention(png, jpeg, jpg)", http.StatusForbidden, "error", nil)
			c.JSON(http.StatusForbidden, response)
			return
		}

		// buat angka acaka untuk me-rename name file
		randomNumb, _ := rand.Int(rand.Reader, big.NewInt(999999999999))
		// sambung nama file dengan angka acak tersebut, lalu set path lokasinya
		path := fmt.Sprintf("images/products/%v-%d-%s", newProduct.ID, randomNumb, file.Filename)

		// mapping datanya
		formatProductImages := product.CreateProductImagesInput{}
		formatProductImages.IsPrimary = "false"
		formatProductImages.Name = path
		formatProductImages.ProductId = newProduct.ID

		// simpan ke database
		save, err := h.service.SaveUploadProductImages(formatProductImages)
		if err != nil || !save {
			response := helper.APIResponse("Failed create product image", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// upload gambarnya satu persatu
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"is_upload": false}
			response := helper.APIResponse("Failed upload file image", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	response := helper.APIResponse("Product created", http.StatusOK, "success", product.FormatterProduct(newProduct))
	c.JSON(http.StatusOK, response)
}
