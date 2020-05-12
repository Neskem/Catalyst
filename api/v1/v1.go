package v1

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/footprint", footprint)
	}
}
