package v1

import (
	"net/http"

	"github.com/afdolriski/restful-test/api/v1/image"
	"github.com/afdolriski/restful-test/api/v1/users"
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", pong)
		users.ApplyRoutes(v1)
		image.ApplyRoutes(v1)
	}
}
