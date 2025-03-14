package controllers

import (
	"techincal-test/responses"
	"techincal-test/services"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	message, status, err := services.AddProductService(c)
	if err != nil {
		responses.AbortResponse(c, status, err.Error())
		return
	}
	responses.CommonResponse(c, status, message)
}

func GetAllProduct(c *gin.Context) {
	products, status, err := services.GetAllProductService(c)
	if err != nil {
		responses.AbortResponse(c, status, err.Error())
		return
	}
	responses.CommonResponse(c, status, products)
}

func GetProduct(c *gin.Context) {
	product, status, err := services.GetProductService(c)
	if err != nil {
		responses.AbortResponse(c, status, err.Error())
		return
	}

	responses.CommonResponse(c, status, product)
}

func UpdateProduct(c *gin.Context) {
	message, status, err := services.UpdateProductService(c)
	if err != nil {
		responses.AbortResponse(c, status, err.Error())
		return
	}
	responses.CommonResponse(c, status, message)
}

func DeleteProduct(c *gin.Context) {
	message, status, err := services.DeleteProductService(c)
	if err != nil {
		responses.AbortResponse(c, status, err.Error())
		return
	}

	responses.CommonResponse(c, status, message)
}
