package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FilterLowStock(c *gin.Context, query *gorm.DB) {
	lowStock := c.DefaultQuery("lowstock", "0")
	if lowStock != "0" {
		query = query.Where("quantity < ?", "50")
	}
}
