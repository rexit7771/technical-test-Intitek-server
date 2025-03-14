package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"
	"techincal-test/helpers"
	"techincal-test/responses"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			responses.AbortResponse(c, http.StatusUnauthorized, "Invalid Token")
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			responses.AbortResponse(c, http.StatusUnauthorized, "Invalid Token")
			return
		}

		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
		var secret_key = os.Getenv("JWT_SECRET_KEY")
		tokenPart := parts[1]
		claims := &helpers.Claims{}
		token, err := jwt.ParseWithClaims(tokenPart, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret_key), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				responses.AbortResponse(c, http.StatusUnauthorized, "Invalid Token Signature")
				return
			}

			responses.AbortResponse(c, http.StatusUnauthorized, "Invalid Token")
			return
		}

		if !token.Valid {
			responses.AbortResponse(c, http.StatusUnauthorized, "invalid Token")
			return
		}
		c.Set("userID", claims.ID)
		c.Next()
	}
}
