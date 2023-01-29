package main

import (
	"LuxSpace/app/v1/category"
	"LuxSpace/app/v1/courir"
	"LuxSpace/configs"
	"LuxSpace/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	courirRepository := courir.NewRepository(db)
	courirService := courir.NewService(courirRepository)
	courirHandler := handler.NewCourirHandler((courirService))

	// fmt.Println(courirService.GetCourirs())

	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	// Category
	apiV1.GET("/category", categoryHandler.GetCategorys)
	apiV1.GET("/category/:id", categoryHandler.GetDetailCategory)
	apiV1.POST("/category", categoryHandler.CreateCategory)
	apiV1.PUT("/category/:id", categoryHandler.UpdateCategory)
	apiV1.DELETE("/category/:id", categoryHandler.DeleteCategory)
	// Courir
	apiV1.GET("/courirs", courirHandler.GetCourirs)
	apiV1.GET("/courir/:id", courirHandler.GetDetailCourir)

	router.Run()
}
