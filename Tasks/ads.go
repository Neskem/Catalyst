package Tasks

import (
	"catalyst.Go/common"
	j "catalyst.Go/database/json"
	"catalyst.Go/kafka"
	"encoding/json"
	guuid "github.com/google/uuid"
	"os"
	"strconv"
)

func UrlAppendAdsInfo(requestJson j.FootPrintRequestBody) bool {
	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	producer := kafka.Producer(kafkaAddress)
	if requestJson.Ads != nil {
		if len(requestJson.Ads) == 0 {
			requestJson.AdsBatchId = ""
			requestJson.AdsBatchSize = "0"
		} else {
			adsBatchId := guuid.New().String()
			requestJson.AdsBatchSize = strconv.Itoa(len(requestJson.Ads))
			for i := range requestJson.Ads {
				tempJson := requestJson
				tempJson.AdsBatchId = adsBatchId
				tempJson = GetRowKey(tempJson)
				rowKey := tempJson.HbaseRowKey
				message := common.Struct2Map(tempJson)
				delete(message, "ads")
				for key, value := range requestJson.Ads[i] {
					message["ads_" + key] = value
				}
				for key, value := range message {
					if value == nil || value == ""{
						delete(message, key)
					}
				}
				messageJson, _ := json.Marshal(message)
				kafka.SendMessage(producer, KafkaTopicAds, messageJson, rowKey)
			}
		}
	} else {
		requestJson.AdsBatchId = ""
		requestJson.AdsBatchSize = "0"
		requestJson = GetRowKey(requestJson)
		message, err := json.Marshal(requestJson)
		if err != nil {
			return false
		}
		kafka.SendMessage(producer, KafkaTopicAds, message, requestJson.HbaseRowKey)
	}
	return true
}
