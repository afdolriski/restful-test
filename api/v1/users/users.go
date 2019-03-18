package users

import (
	"github.com/gin-gonic/gin"

	controllers "github.com/afdolriski/restful-test/controllers/users"
)

// ApplyRoutes to gin engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/users")
	{
		posts.GET("/", controllers.Index)
		posts.POST("/", controllers.Create)
		posts.GET("/:id", controllers.Browse)
		posts.PUT("/:id", controllers.Update)
		posts.DELETE("/:id", controllers.Delete)
	}
}
