package image

import (
	"github.com/gin-gonic/gin"

	controllers "github.com/afdolriski/restful-test/controllers/image"
)

// ApplyRoutes to gin engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/fetch")
	{
		posts.GET("/", controllers.Fetch)
	}
}
