package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rubenbuelvas/garage-be/src/api/adapter/controller"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", controller.Ping())

	// Product routes
	productGroup := router.Group("/products")
	{
		// CRUD operations
		productGroup.POST("/", controller.CreateProduct)
		productGroup.GET("/", controller.GetProducts)
		productGroup.GET("/:id", controller.GetProductByID)
		productGroup.PUT("/:id", controller.UpdateProduct)
		productGroup.DELETE("/:id", controller.DeleteProduct)

		// Buy
		productGroup.POST("/:id/buy", controller.BuyProduct)

		// Search
		productGroup.GET("/search", controller.SearchProductsByQuery)
	}

	// User routes
	userGroup := router.Group("/admin")
	{
		userGroup.POST("/", controller.CreateUser)
		userGroup.GET("/", controller.GetUsers)
		userGroup.GET("/:id", controller.GetUserByID)
		userGroup.PUT("/:id", controller.UpdateUser)
		userGroup.DELETE("/:id", controller.DeleteUser)
	}
}
