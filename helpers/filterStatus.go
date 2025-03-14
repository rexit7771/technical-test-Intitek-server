package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FilterStatus(c *gin.Context, query *gorm.DB) {
	status := c.DefaultQuery("status", "")
	if status != "" {
		statusQuery := "%" + status + "%"
		query = query.Where("status LIKE ?", statusQuery)
	}
}
