package responses

import "github.com/gin-gonic/gin"

func AbortResponse(c *gin.Context, status int, message interface{}) {
	switch message {
	case "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)":
		message = "Email has been used"
	}
	c.AbortWithStatusJSON(status, gin.H{"message": message})
}
