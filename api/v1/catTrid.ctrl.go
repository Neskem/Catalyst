package v1

import (
	"catalyst.Go/common"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

func catTrid(c *gin.Context) {
	var ptJwtValue string
	var catTridValue string
	ptJwt, err := c.Request.Cookie("pt_jwt")
	if err != nil {
		ptJwtValue = ""
	} else {
		ptJwtValue = ptJwt.Value
	}
	catTrid, err := c.Request.Cookie("cat_trid")
	if err != nil {
		now := strconv.FormatInt(time.Now().UnixNano(), 10)
		catTridValue = guuid.New().String() + now[:10] + "." + now[10:17]
	} else {
		catTridValue = catTrid.Value
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:       "cat_trid",
		Value:      catTridValue,
		Path:       "/",
		Domain:     ".breaktime.com.tw",
		Expires:    time.Now().AddDate(2, 0, 0),
		RawExpires: "",
		MaxAge:     60 * 60 * 24 * 365 * 2,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})
	c.JSON(200, common.JSON{
		"cat_trid": catTridValue,
		"pt_jwt": ptJwtValue,
	})
}
