package v1

import (
	"catalyst.Go/Tasks"
	"catalyst.Go/common"
	j "catalyst.Go/database/json"
	"catalyst.Go/kafka"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"github.com/ua-parser/uap-go/uaparser"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var KafkaTopicFootprint = "footprint"
var KafkaTopicProfileYahoo = "profile_oath"

func footprint(c *gin.Context) {
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

	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	producer := kafka.Producer(kafkaAddress)
	message, err := json.Marshal(requestBody)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	kafka.SendMessage(producer, KafkaTopicFootprint, message, requestBody.HbaseRowKey)
	c.JSON(200, common.JSON{
		"page_id": requestBody.PageId,
	})
}

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

func profileOath(c *gin.Context) {
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
	requestBody = Tasks.UrlAppendUserInfo(requestBody)

	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	producer := kafka.Producer(kafkaAddress)
	message, err := json.Marshal(requestBody)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	kafka.SendMessage(producer, KafkaTopicProfileYahoo, message, requestBody.HbaseRowKey)
	c.AbortWithStatus(200)
}
