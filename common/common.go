package common

import (
	"crypto/sha1"
	"encoding/hex"
	"net"
	"net/url"
	"strings"
)

type JSON = map[string]interface{}


func ValidateIpAddress(ip string) bool {
	ipValidate := net.ParseIP(ip)
	if ipValidate != nil && (ipValidate.To4() != nil || ipValidate.To16() != nil) {
		return true
	}
	return false
}

func GetPageID(unParsedUrl string) string {
	unParsedUrl = strings.ReplaceAll(unParsedUrl, "#.*", "")
	parsedUrl, err := url.Parse(unParsedUrl)
	if err != nil {
		panic(err)
	}
	netLoc := parsedUrl.Hostname()
	if len(parsedUrl.Port()) > 0 {
		netLoc = netLoc + parsedUrl.Port()
	}
	var text string
	if len(parsedUrl.Query()) > 0 {
		text = netLoc + parsedUrl.Path + parsedUrl.RawQuery
	} else {
		text = netLoc + parsedUrl.Path
	}
	hashCode := sha1.New()
	hashCode.Write([]byte(text))
	return hex.EncodeToString(hashCode.Sum(nil))
}
