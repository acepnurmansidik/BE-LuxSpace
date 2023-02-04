package main

import (
	"LuxSpace/app/v1/address"
	"LuxSpace/app/v1/category"
	"LuxSpace/app/v1/courir"
	"LuxSpace/app/v1/user"
	"LuxSpace/auth"
	"LuxSpace/configs"
	"LuxSpace/handler"
	"LuxSpace/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := configs.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Auth
	authService := auth.NewService()

	// V1 - Category
	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	// V1 - Courir
	courirRepository := courir.NewRepository(db)
	courirService := courir.NewService(courirRepository)
	courirHandler := handler.NewCourirHandler((courirService))
	// V1 - User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)
	// V1 - Address
	addressRepository := address.NewRepository(db)
	addressService := address.NewService(addressRepository)
	addressHandler := handler.NewAddressHandler(addressService)

	authMiddleware := middleware.NewAuthMiddleware(authService, userService)

	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	// added auth middleware
	apiV1.Use(authMiddleware.AuthMiddleware())
	// Category
	apiV1.GET("/category", categoryHandler.GetCategorys)
	apiV1.GET("/category/:id", categoryHandler.GetDetailCategory)
	apiV1.POST("/category", categoryHandler.CreateCategory)
	apiV1.PUT("/category/:id", categoryHandler.UpdateCategory)
	apiV1.DELETE("/category/:id", categoryHandler.DeleteCategory)
	// Courir
	apiV1.GET("/courirs", courirHandler.GetCourirs)
	apiV1.GET("/courir/:id", courirHandler.GetDetailCourir)
	apiV1.POST("/courir", courirHandler.CreateCourir)
	apiV1.PUT("/courir/:id", courirHandler.UpdateCourir)
	apiV1.DELETE("/courir/:id", courirHandler.DeleteCourir)
	// User
	apiV1.POST("/register", userHandler.RegisterUser)
	apiV1.POST("/login", userHandler.LoginUser)
	apiV1.POST("/activate/:otp", userHandler.ActivateUser)
	// Address
	apiV1.GET("/address-list", addressHandler.GetAllAddress)
	apiV1.POST("/address", addressHandler.CreateAddress)
	apiV1.PUT("/address/:id", addressHandler.UpdateAddress)

	router.Run()
}
