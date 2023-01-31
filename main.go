package main

import (
	"LuxSpace/app/v1/category"
	"LuxSpace/app/v1/courir"
	"LuxSpace/app/v1/user"
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
	userHandler := handler.NewUserHandler(userService)

	// user := user.CreateUserInput{
	// 	Username: "acep",
	// 	Email:    "acep@gmail.com",
	// 	Password: "12345",
	// }
	// newUser, _ := userService.Register(user)
	// fmt.Println(newUser)

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
	apiV1.POST("/courir", courirHandler.CreateCourir)
	apiV1.PUT("/courir/:id", courirHandler.UpdateCourir)
	apiV1.DELETE("/courir/:id", courirHandler.DeleteCourir)
	// User
	apiV1.POST("/register", userHandler.RegisterUser)

	router.Run()
}
