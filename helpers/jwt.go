package helpers

import (
	"log"
	"os"
	"techincal-test/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func SignPayload(payload structs.User) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	var secret_key = os.Getenv("JWT_SECRET_KEY")

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		ID: payload.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
