package v1

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/footprint", footprint)
		v1.POST("/profile_oath", profileOath)
		v1.POST("/wifi", wifi)
		v1.POST("/open_id", openId)
		v1.GET("/cat_trid", catTrid)
	}
}
