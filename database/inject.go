package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Inject db to gin context
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
