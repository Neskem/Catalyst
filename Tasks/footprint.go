package Tasks

import (
	"catalyst.Go/common"
	"catalyst.Go/database/json"
	guuid "github.com/google/uuid"
	"github.com/ua-parser/uap-go/uaparser"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func UrlExtractionJudgement(requestJson json.FootPrintRequestBody)(json.FootPrintRequestBody, bool){
	requestJson, rtn := UrlExtractionPageId(requestJson)
	if rtn == false {
		return requestJson, false
	}
	UrlExtractionUrl(requestJson)
	UrlExtractionUrlOg(requestJson)
	UrlExtractionUrlCanonical(requestJson)
	UrlExtractionReferrer(requestJson)
	return requestJson, true
}

func UrlExtractionPageId(requestJson json.FootPrintRequestBody)(json.FootPrintRequestBody, bool){
	if requestJson.UrlPageId != "" {
		tmpDecode, err := url.QueryUnescape(requestJson.UrlPageId)
		if err != nil {
			return requestJson, false
		}
		urlObject, err := url.Parse(tmpDecode)
		if err != nil {
			return requestJson, false
		}
		if urlObject.Host == "" {
			return requestJson, false
		} else {
			requestJson.UrlPageIdDecode = tmpDecode
			requestJson.UrlPageIdHostname = urlObject.Host
			return requestJson, true
		}
	}
	return requestJson, false
}

func UrlExtractionUrl(requestJson json.FootPrintRequestBody)(json.FootPrintRequestBody, bool){
	if requestJson.Url != "" {
		tmpDecode, err := url.QueryUnescape(requestJson.Url)
		if err != nil {
			return requestJson, false
		}
		urlObject, err := url.Parse(tmpDecode)
		if err != nil {
			return requestJson, false
		}
		if urlObject.Host == "" {
			return requestJson, false
		} else {
			requestJson.UrlDecode = tmpDecode
			requestJson.UrlHostname = urlObject.Host
		}
	}
	return requestJson, false
}

func UrlExtractionUrlOg(requestJson json.FootPrintRequestBody)(json.FootPrintRequestBody, bool){
	if requestJson.UrlOg != "" {
		tmpDecode, err := url.QueryUnescape(requestJson.UrlOg)
		if err != nil {
			return requestJson, false
		}
		urlObject, err := url.Parse(tmpDecode)
		if err != nil {
			return requestJson, false
		}
		if urlObject.Host == "" {
			return requestJson, false
		} else {
			requestJson.UrlOgDecode = tmpDecode
			requestJson.UrlOgHostname = urlObject.Host
		}
	}
	return requestJson, false
}

func UrlExtractionUrlCanonical(requestJson json.FootPrintRequestBody)(json.FootPrintRequestBody, bool){
	if requestJson.UrlCanonical != "" {
		tmpDecode, err := url.QueryUnescape(requestJson.UrlCanonical)
		if err != nil {
			return requestJson, false
		}
		urlObject, err := url.Parse(tmpDecode)
		if err != nil {
			return requestJson, false
		}
		if urlObject.Host == "" {
			return requestJson, false
		} else {
			requestJson.UrlCanonicalDecode = tmpDecode
			requestJson.UrlCanonicalHostname = urlObject.Host
		}
	}
	return requestJson, false
}

func UrlExtractionReferrer(requestJson json.FootPrintRequestBody)(json.FootPrintRequestBody, bool){
	if requestJson.Referrer != "" {
		tmpDecode, err := url.QueryUnescape(requestJson.Referrer)
		if err != nil {
			return requestJson, false
		}
		urlObject, err := url.Parse(tmpDecode)
		if err != nil {
			return requestJson, false
		}
		if urlObject.Host == "" {
			return requestJson, false
		} else {
			requestJson.ReferrerDecode = tmpDecode
			requestJson.ReferrerHostname = urlObject.Host
		}
	}
	return requestJson, false
}

func UserAgentExtraction(requestJson json.FootPrintRequestBody, client *uaparser.Client) json.FootPrintRequestBody {
	requestJson.UaBrowserFamily = client.UserAgent.Family
	requestJson.UaBrowserVersionMajor = client.UserAgent.Major
	requestJson.UaBrowserVersionMinor = client.UserAgent.Minor
	requestJson.UaBrowserVersionBuild = client.UserAgent.Patch
	requestJson.UaBrowserVersionString = client.UserAgent.ToVersionString()

	requestJson.UaOsFamily = client.Os.Family
	requestJson.UaOsVersionMajor = client.Os.Major
	requestJson.UaOsVersionMinor = client.Os.Minor
	requestJson.UaOsVersionBuild = client.Os.Patch
	requestJson.UaOsVersionString = client.Os.ToVersionString()

	requestJson.UaDeviceFamily = client.Device.Family
	requestJson.UaDeviceBrand = client.Device.Brand
	requestJson.UaDeviceModel = client.Device.Model

	requestJson.UaIsMobile = common.IsMobile(client)
	requestJson.UaIsTablet = common.IsTablet(client)
	requestJson.UaIsTouchCapable = common.IsTouchCapable(client)
	requestJson.UaIsPc = common.IsPc(client)
	requestJson.UaIsBot = common.IsBot(client)

	return requestJson
}

func UrlExtractPageId(requestJson json.FootPrintRequestBody) json.FootPrintRequestBody {
	if requestJson.UrlPageIdDecode != "" {
		requestJson.PageId = common.GetPageID(requestJson.UrlPageIdDecode)
	}
	return requestJson
}

func UrlAppendDatetime(requestJson json.FootPrintRequestBody) json.FootPrintRequestBody {
	loc, _ := time.LoadLocation("Asia/Taipei")
	dateTime := time.Now().In(loc)
	requestJson.CreationTime = strconv.Itoa(dateTime.Year()) + "-" + dateTime.Month().String() + "-" + strconv.Itoa(dateTime.Day()) + "T" + strconv.Itoa(dateTime.Hour()) + ":" + strconv.Itoa(dateTime.Minute()) + ":" + strconv.Itoa(dateTime.Second()) + " +0800"
	return requestJson
}

func GetRowKey(requestJson json.FootPrintRequestBody) json.FootPrintRequestBody {
	prefixAlpha := [7]string{"a", "b", "c", "d", "e", "f", "g"}
	choose := prefixAlpha[rand.Intn(len(prefixAlpha))]
	id := strings.ReplaceAll(guuid.New().String(), "-", "", )
	tmpCt := requestJson.CreationTime[:19]
	requestJson.HbaseRowKey = choose + "_" + tmpCt + "_" + id
	return requestJson
}
