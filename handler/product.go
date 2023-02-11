package handler

import (
	"LuxSpace/app/v1/merchant"
	"LuxSpace/app/v1/product"
	"LuxSpace/app/v1/user"
	"LuxSpace/helper"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service  product.Service
	merchant merchant.Service
}

func NewProductHandler(service product.Service, merchant merchant.Service) *productHandler {
	return &productHandler{service, merchant}
}

func (h *productHandler) CreateProductMerchant(c *gin.Context) {
	var inputDataProduct product.CreateProductInput
	err := c.ShouldBind(&inputDataProduct)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil ID user nya
	value := c.MustGet("currentUser")
	userID := value.(user.FormatUserHeader)
	// ambil id merchant berdasarkan user yang sudja daftar
	merchantID, err := h.merchant.GetMerchantByUserID(userID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch data product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputDataProduct.MerchantID = merchantID.ID

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
		formatProductImages.IsDelete = "false"
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

	response := helper.APIResponse("Product created", http.StatusOK, "success", product.FormatterMerchantProducts(newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetAllMerchantProduct(c *gin.Context) {
	// ambil id usernya di JWT
	value := c.MustGet("currentUser")
	userID := value.(user.FormatUserHeader)

	// cari merchant id
	merchantRegister, err := h.merchant.GetMerchantByUserID(userID.ID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data merchant", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.service.GetAllMerchantProduct(merchantRegister.ID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get products", http.StatusOK, "success", product.FormatterMerchantProductsList(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	var inputID product.ProductDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// delete datanya
	newProduct, err := h.service.DeleteProduct(inputID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product deleted", http.StatusOK, "success", product.FormatterMerchantProducts(newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputData product.CreateProductInput
	err := c.ShouldBind(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed fetch data products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil id productnya
	var inputID product.ProductDetailInput
	err = c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data uri products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// update datanya
	newProduct, err := h.service.UpdateProductMerchant(inputData, inputID)
	if err != nil {
		response := helper.APIResponse("Failed fetch data products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data updated", http.StatusOK, "succcess", product.FormatterMerchantProducts(newProduct))
	c.JSON(http.StatusOK, response)
}
