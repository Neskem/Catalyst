package v1

import (
	"catalyst.Go/Tasks"
	"catalyst.Go/common"
	j "catalyst.Go/database/json"
	"catalyst.Go/kafka"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
	"log"
	"os"
	"strings"
)

func footprint(c *gin.Context) {
	var requestBody j.FootPrintRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}
	requestBody, rtn := Tasks.UrlExtractionJudgement(requestBody)
	if rtn == false {
		c.AbortWithStatus(500)
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

	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	producer := kafka.Producer(kafkaAddress)
	message, err := json.Marshal(requestBody)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	kafka.SendMessage(producer, "sync_topic", message, "a_2018-07-06T10:12:02.1234567")
	c.AbortWithStatus(200)
	return
}
