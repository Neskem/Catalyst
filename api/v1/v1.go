package v1

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/footprint", footprint)
		v1.POST("/profile_oath", profileOath)
		v1.POST("/wifi", wifi)
		v1.POST("/open_id", openId)
		v1.POST("/ads", ads)
		v1.POST("/conversion",  conversion)
		v1.POST("/highlighted_text", highlightedText)
		v1.POST("/openlink", openLink)
		v1.POST("/session_stay", sessionStay)
		v1.GET("/cat_trid", catTrid)
	}
}
