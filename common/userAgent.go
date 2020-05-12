package common

import (
	"github.com/ua-parser/uap-go/uaparser"
	"strings"
)

var MobileDeviceFamilies = [6] string{"iPhone", "iPod", "Generic Smartphone", "Generic Feature Phone", "PlayStation Vita", "iOS-Device"}
var MobileBrowserFamilies = [2] string {"Opera Mobile", "Opera Mini"}
var TabletDeviceFamilies = [9] string {"iPad", "BlackBerry Playbook", "Blackberry Playbook", "Kindle", "Kindle Fire", "Kindle Fire HD", "Galaxy Tab", "Xoom", "Dell Streak"}
var MobileOsFamilies = [7] string {"Windows Phone", "Windows Phone OS", "Symbian OS",  "Bada", "Windows CE", "Windows Mobile", "Maemo"}
var TouchCapableOsFamilies = [9] string {"iOS", "Android", "Windows Phone", "Windows Phone OS", "Windows RT", "Windows CE", "Windows Mobile", "Firefox OS", "MeeGo"}
var TouchCapableDeviceFamilies = [3] string {"BlackBerry Playbook", "Blackberry Playbook", "Kindle Fire"}
var PcOsFamilies = [4] string {"Windows 95", "Windows 98", "Windows ME", "Solaris"}
var EmailProgramFamilies = [16] string {"Outlook", "Windows Live Mail", "AirMail", "Apple Mail", "Outlook", "Thunderbird", "Lightning", "ThunderBrowse", "Windows Live Mail",
	"The Bat!", "Lotus Notes", "IBM Notes", "Barca", "MailBar", "kmail2", "YahooMobileMail"}

func IsMobile(client *uaparser.Client)bool{
	// First check for mobile device and mobile browser families
	for _, device := range MobileDeviceFamilies {
		if client.Device.Family == device {
			return true
		}
	}
	for _, device := range MobileBrowserFamilies {
		if client.UserAgent.Family == device {
			return true
		}
	}
	// Device is considered Mobile OS is Android and not tablet
	// This is not fool proof but would have to suffice for now
	if (client.Os.Family == "Android" || client.Os.Family == "Firefox OS") && !IsTablet(client){
		return true
	}
	if client.Os.Family == "BlackBerry OS" && client.Device.Family != "Blackberry Playbook" {
		return true
	}
	for _, device := range MobileOsFamilies {
		if client.Os.Family == device {
			return true
		}
	}
	// TODO: remove after https://github.com/tobie/ua-parser/issues/126 is closed
	if strings.Contains(client.UserAgent.ToString(), "J2ME") || strings.Contains(client.UserAgent.ToString(), "MIDP") {
		return true
	}
	// This is here mainly to detect Google's Mobile Spider
	if strings.Contains(client.UserAgent.ToString(), "iPhone") {
		return true
	}
	if strings.Contains(client.UserAgent.ToString(), "Googlebot-Mobile") {
		return true
	}
	// Mobile Spiders should be identified as mobile
	if client.Device.Family == "Spider" && strings.Contains(client.UserAgent.ToString(), "Mobile") {
		return true
	}
	if strings.Contains(client.UserAgent.ToString(), "NokiaBrowser") && strings.Contains(client.UserAgent.ToString(), "Mobile") {
		return true
	}
	return false
}

func IsTouchCapable(client *uaparser.Client)bool{
	for _, device := range TouchCapableOsFamilies {
		if client.Os.Family == device {
			return true
		}
	}
	for _, device := range TouchCapableDeviceFamilies {
		if client.Device.Family == device {
			return true
		}
	}
	if strings.HasPrefix(client.Os.Family, "Windows 8") && strings.Contains(client.UserAgent.ToString(), "Touch") {
		return true
	}
	if strings.Contains(client.Os.Family, "BlackBerry") && IsBlackBerryTouchCapableDevice(client) {
		return true
	}
	return false
}

func IsPc(client *uaparser.Client)bool{
	// Returns True for "PC" devices (Windows, Mac and Linux)
	if strings.Contains(client.UserAgent.ToString(), "Windows NT") {
		return true
	}
	for _, device := range PcOsFamilies {
		if client.Os.Family == device {
			return true
		}
	}
	// TODO: remove after https://github.com/tobie/ua-parser/issues/127 is closed
	if client.Os.Family == "Mac OS X" && !strings.Contains(client.UserAgent.ToString(), "Silk") {
		return true
	}
	// Maemo has 'Linux' and 'X11' in UA, but it is not for PC
	if strings.Contains(client.UserAgent.ToString(), "Maemo") {
		return false
	}
	if strings.Contains(client.Os.Family, "Chrome OS") {
		return true
	}
	if strings.Contains(client.UserAgent.ToString(), "Linux") && strings.Contains(client.UserAgent.ToString(), "X11") {
		return true
	}
	return false
}

func IsBot(client *uaparser.Client)bool{
	if client.Device.Family == "Spider" {
		return true
	}
	return false
}
func IsEmailClient(client *uaparser.Client)bool{
	for _, device := range EmailProgramFamilies{
		if client.UserAgent.Family == device {
			return true
		}
	}
	return false
}

func IsTablet(client *uaparser.Client)bool{
	for _, device := range TabletDeviceFamilies {
		if client.Device.Family == device {
			return true
		}
	}
	if client.Os.Family == "Android" && IsAndroidTablet(client) {
		return true
	}
	if strings.HasPrefix(client.Os.Family, "Windows RT") {
		return true
	}
	if client.Os.Family == "Firefox OS" && !strings.Contains(client.UserAgent.Family, "Mobile"){
		return true
	}
	return false
}

func IsAndroidTablet(client *uaparser.Client)bool{
	if strings.Contains(client.UserAgent.ToString(), "Mobile Safari") && client.UserAgent.Family != "Firefox Mobile" {
		return true
	}
	return false
}

func IsBlackBerryTouchCapableDevice(client *uaparser.Client)bool{
	// A helper to determine whether a BB phone has touch capabilities
	// Blackberry Bold Touch series begins with 99XX
	if strings.Contains(client.Device.Family, "Blackberry 99") {
		return true
	}
	if strings.Contains(client.Device.Family, "Blackberry 95") {
		return true
	}
	return false
}
