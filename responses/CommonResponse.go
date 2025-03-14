package responses

import "github.com/gin-gonic/gin"

func CommonResponse(c *gin.Context, status int, message interface{}) {
	c.JSON(status, gin.H{"result": message})
}
