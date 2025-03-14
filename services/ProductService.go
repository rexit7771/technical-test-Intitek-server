package services

import (
	"errors"
	"net/http"
	"strconv"
	"techincal-test/database"
	"techincal-test/helpers"
	"techincal-test/structs"

	"github.com/gin-gonic/gin"
)

func AddProductService(c *gin.Context) (message string, status int, err error) {
	var newProduct structs.Product
	if err = c.ShouldBindJSON(&newProduct); err != nil {
		return "", http.StatusInternalServerError, err
	}

	if err := newProduct.Validate(); err != nil {
		return "", http.StatusBadRequest, err
	}

	if err := database.DB.Create(&newProduct).Error; err != nil {
		return "", http.StatusInternalServerError, err
	}

	return "New Product has been added", http.StatusOK, nil
}

func GetAllProductService(c *gin.Context) (products []structs.Product, status int, err error) {
	query := database.DB.Model(&structs.Product{})
	if query.Error != nil {
		return products, http.StatusInternalServerError, query.Error
	}
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return products, http.StatusInternalServerError, err
	}
	if count == 0 {
		return products, http.StatusNotFound, errors.New("no products found")
	}

	helpers.FilterStatus(c, query)
	helpers.FilterLowStock(c, query)
	query.Find(&products)
	return products, http.StatusOK, nil
}

func GetProductService(c *gin.Context) (product structs.Product, status int, err error) {
	id := c.Param("id")
	if err := database.DB.First(&product, id).Error; err != nil {
		return product, http.StatusInternalServerError, err
	}
	if product.ID == 0 {
		return product, http.StatusNotFound, errors.New("product is not found")
	}
	return product, http.StatusOK, nil
}

func UpdateProductService(c *gin.Context) (message string, status int, err error) {
	product, status, err := GetProductService(c)
	if err != nil {
		return "", status, err
	}

	var productUpdated structs.Product
	err = c.ShouldBindJSON(&productUpdated)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	err = productUpdated.Validate()
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	err = database.DB.Model(&product).Updates(productUpdated).Error
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	productID := strconv.FormatUint(uint64(product.ID), 10)
	message = "Product with ID " + productID + " has been updated"
	return message, http.StatusOK, nil
}

func DeleteProductService(c *gin.Context) (message string, status int, err error) {
	product, status, err := GetProductService(c)
	if err != nil {
		return "", status, err
	}

	err = database.DB.Delete(&product).Error
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	productID := strconv.FormatUint(uint64(product.ID), 10)
	message = "Product with ID " + productID + " has been deleted"
	return message, http.StatusOK, nil
}
