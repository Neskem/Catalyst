package v1

import (
	"catalyst.Go/Tasks"
	"catalyst.Go/common"
	j "catalyst.Go/database/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
	"log"
	"strings"
)

func ads(c *gin.Context) {
	var requestBody j.FootPrintRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.JSON(500, common.JSON{
			"message": "Required fields are missing.",
			"payload": requestBody,
		})
		return
	}
	requestBody, rtn := Tasks.UrlExtractionJudgement(requestBody)
	if rtn == false {
		c.JSON(500, common.JSON{
			"message": "Required fields are missing.",
			"payload": requestBody,
		})
		return
	}

	parser, err := uaparser.New("./regexes.yaml")
	if err != nil {
		log.Fatal(err)
	}

	requestBody.Ua = c.Request.Header.Get("User-Agent")

	client := parser.Parse(requestBody.Ua)
	Tasks.UserAgentExtraction(requestBody, client)

	forwardedIp := strings.ToLower(strings.Trim(c.Request.Header.Get("X-Forwarded-For"), " "))
	if common.ValidateIpAddress(forwardedIp) {
		requestBody.Ip = forwardedIp
		requestBody.IpXForwaredFor = forwardedIp
	}

	realIp := strings.ToLower(strings.Trim(c.Request.Header.Get("X-Real-Ip"), " "))
	if common.ValidateIpAddress(realIp) {
		if requestBody.Ip == "" {
			requestBody.Ip = realIp
		}
		requestBody.IpXRealIp = realIp
	}
	requestBody = Tasks.UrlExtractPageId(requestBody)
	requestBody = Tasks.UrlAppendDatetime(requestBody)
	requestBody = Tasks.GetRowKey(requestBody)
	sendKafkaStatus := Tasks.UrlAppendAdsInfo(requestBody)
	if sendKafkaStatus {
		c.AbortWithStatus(200)
		return
	} else {
		c.AbortWithStatus(400)
		return
	}
}
