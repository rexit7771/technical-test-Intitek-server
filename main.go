package main

import (
	"os"
	"techincal-test/controllers"
	"techincal-test/database"
	"techincal-test/middlewares"
	"techincal-test/routes"
	"techincal-test/seeders"
	"techincal-test/structs"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&structs.Product{})
	database.DB.AutoMigrate(&structs.User{})
	seeders.SeedUsers()
	seeders.SeedProducts()

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.POST("/login", controllers.Login)
	routes.ProductRoutes(router, database.DB)
	router.Run(":" + os.Getenv("PORT"))
}
