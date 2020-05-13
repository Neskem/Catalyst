package Tasks

import (
	j "catalyst.Go/database/json"
	"strconv"
)

func UrlAppendSessionStay(requestJson j.FootPrintRequestBody) j.FootPrintRequestBody {
	if SessionStay, err := strconv.Atoi(requestJson.SessionStay); err == nil {
		requestJson.SessionStayInt = SessionStay
	}
	return requestJson
}
