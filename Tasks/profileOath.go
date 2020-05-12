package Tasks

import (
	j "catalyst.Go/database/json"
	"strconv"
	"strings"
)

func UrlAppendUserInfo(requestJson j.FootPrintRequestBody) j.FootPrintRequestBody {
	if UserAge, err := strconv.Atoi(requestJson.UserAge); err == nil {
		requestJson.UserAgeInt = UserAge
	}
	if requestJson.UserGender != "" {
		requestJson.UserGender = strings.ToUpper(requestJson.UserGender)
	}
	if requestJson.UserCountry != "" {
		requestJson.UserCountry = strings.ToUpper(requestJson.UserCountry)
	}
	return requestJson
}
