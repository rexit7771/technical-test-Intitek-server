package services

import (
	"errors"
	"net/http"
	"techincal-test/database"
	"techincal-test/helpers"
	"techincal-test/structs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(c *gin.Context) (token string, status int, err error) {
	var user structs.User
	if err = c.ShouldBindJSON(&user); err != nil {
		return "", http.StatusInternalServerError, err
	}

	if err = user.Validate(); err != nil {
		return "", http.StatusBadRequest, err
	}

	var userDB structs.User
	tx := database.DB.Where("email = ?", user.Email).First(&userDB)
	if tx.Error != nil {
		return "", http.StatusUnauthorized, errors.New("Invalid Email/Password")
	}

	bcryptResult := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if bcryptResult != nil {
		return "", http.StatusUnauthorized, errors.New("Invalid Email/Password")
	}

	token, tokenError := helpers.SignPayload(userDB)
	if tokenError != nil {
		return "", http.StatusInternalServerError, tokenError
	}

	return token, http.StatusOK, nil
}
