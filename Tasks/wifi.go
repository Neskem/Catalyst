package Tasks

import (
	j "catalyst.Go/database/json"
	"strconv"
	"strings"
)

func UrlAppendWifiInfo(requestJson j.FootPrintRequestBody) j.FootPrintRequestBody {
	if requestJson.WLanUserAge != "" {
		wLanUser := strings.Split(requestJson.WLanUserAge, "-")
		if len(wLanUser) > 0 {
			if WLanUserAgeLowerBound, err := strconv.Atoi(wLanUser[0]); err ==nil {
				requestJson.WLanUserAgeLowerBound = WLanUserAgeLowerBound
			} else {
				if strings.Contains(wLanUser[0], "以上") {
					if WLanUserAgeLowerBound, err := strconv.Atoi(strings.ReplaceAll(wLanUser[0], "以上", "")); err ==nil {
						requestJson.WLanUserAgeLowerBound = WLanUserAgeLowerBound
					}
				}
				if strings.Contains(wLanUser[0], "以下") {
					if WLanUserAgeUpperBound, err := strconv.Atoi(strings.ReplaceAll(wLanUser[0], "以下", "")); err ==nil {
						requestJson.WLanUserAgeUpperBound = WLanUserAgeUpperBound
					}
				}
			}
		}
		if len(wLanUser) > 1 {
			if WLanUserAgeUpperBound, err := strconv.Atoi(wLanUser[1]); err ==nil {
				requestJson.WLanUserAgeUpperBound = WLanUserAgeUpperBound
			}
		}
	}
	return requestJson
}
