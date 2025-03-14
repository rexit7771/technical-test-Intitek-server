package controllers

import (
	"techincal-test/responses"
	"techincal-test/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token, status, err := services.LoginService(c)
	if err != nil {
		responses.AbortResponse(c, status, err.Error())
		return
	}
	responses.TokenResponse(c, token)
}
