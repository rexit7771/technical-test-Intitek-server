package routes

import (
	"techincal-test/controllers"
	"techincal-test/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(router *gin.Engine, db *gorm.DB) {
	productsGroup := router.Group("/products")
	productsGroup.Use(middlewares.Authentication())
	{
		productsGroup.POST("/", controllers.AddProduct)
		productsGroup.GET("/", controllers.GetAllProduct)
		productsGroup.GET("/:id", controllers.GetProduct)
		productsGroup.PUT("/:id", controllers.UpdateProduct)
		productsGroup.DELETE("/:id", controllers.DeleteProduct)
	}
}
