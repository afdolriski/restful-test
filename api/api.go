package api

import (
	v1 "github.com/afdolriski/restful-test/api/v1"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes for api
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}
