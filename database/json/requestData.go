package json

type FootPrintRequestBody struct {
	DeviceType string `json:"device_type,omitempty"`
	Fp string `json:"fp" binding:"required"`
	InfinityBatchId string `json:"infinity_batchid,omitempty"`
	InfinityContentPercentage string `json:"infinity_content_percentage,omitempty"`
	InfinityContentSeqId string `json:"infinity_content_seqid,,omitempty"`
	InfinityUrl string `json:"infinity_url,omitempty"`
	InfinityUrlFirstPage string `json:"infinity_url_firstpage,omitempty"`
	IsInfinity string `json:"is_infinity,omitempty"`
	Referrer string `json:"referrer,omitempty"`
	SessionId string `json:"session_id,omitempty"`
	Spj string `json:"spj,omitempty"`
	TxnId string `json:"txn_id" binding:"required"`
	TxnId2 string `json:"txn_id2,omitempty"`
	Url string `json:"url,omitempty"`
	UrlCanonical string `json:"url_canonical,omitempty"`
	UrlOg string `json:"url_og,omitempty"`
	UrlPageId string `json:"url_pageid" binding:"required"`

	UrlCanonicalDecode string `json:"url_canonical(decode),omitempty"`
	ReferrerDecode string `json:"referrer(decode),omitempty"`
	UrlPageIdDecode string `json:"url_pageid(decode),omitempty"`
	UrlDecode string `json:"url(decode),omitempty"`
	UrlOgDecode string `json:"url_og(decode),omitempty"`
	UrlPageIdHostname string `json:"url_pageid_hostname,omitempty"`
	UrlHostname string `json:"url_hostname,omitempty"`
	UrlCanonicalHostname string `json:"url_canonical_hostname,omitempty"`
	UrlOgHostname string `json:"url_og_hostname,omitempty"`
	ReferrerHostname string `json:"referrer_hostname,omitempty"`

	Ua string `json:"ua,omitempty"`
	UaBrowserFamily string `json:"ua_browser_family,omitempty"`
	UaBrowserVersionMajor string `json:"ua_browser_version_major,omitempty"`
	UaBrowserVersionMinor string `json:"ua_browser_version_minor,omitempty"`
	UaBrowserVersionBuild string `json:"ua_browser_version_build,omitempty"`
	UaBrowserVersionString string `json:"ua_browser_version_string,omitempty"`
	UaOsFamily string `json:"ua_os_family,omitempty"`
	UaOsVersionMajor string `json:"ua_os_version_major,omitempty"`
	UaOsVersionMinor string `json:"ua_os_version_minor,omitempty"`
	UaOsVersionBuild string `json:"ua_os_version_build,omitempty"`
	UaOsVersionString string `json:"ua_os_version_string,omitempty"`
	UaDeviceFamily string `json:"ua_device_family,omitempty"`
	UaDeviceBrand string `json:"ua_device_brand,omitempty"`
	UaDeviceModel string `json:"ua_device_model,omitempty"`
	UaIsMobile bool `json:"ua_is_mobile,omitempty"`
	UaIsTablet bool `json:"ua_is_tablet,omitempty"`
	UaIsTouchCapable bool `json:"ua_is_touch_capable,omitempty"`
	UaIsPc bool `json:"ua_is_pc,omitempty"`
	UaIsBot bool `json:"ua_is_bot,omitempty"`

	Ip string `json:"ip,omitempty"`
	IpXForwaredFor string `json:"ip_X-Forwarded-For,omitempty"`
	IpXRealIp string `json:"ip_X-Real-Ip,omitempty"`

	PageId string `json:"page_id,omitempty"`

	CreationTime string `json:"creation_time,omitempty"`

	HbaseRowKey string `json:"hbase_rowkey,omitempty"`

	// Only use for profile_oath api
	UserAge string `json:"user_age,omitempty"`
	UserAgeInt int `json:"user_age(int),omitempty"`
	UserGender string `json:"user_gender,omitempty"`
	UserCountry string `json:"user_country,omitempty"`

	// Only use for wifi api
	WLanUserAge string `json:"wlan_userage,omitempty"`
	WLanUserAgeLowerBound int `json:"wlan_userage_lower_bound,omitempty"`
	WLanUserAgeUpperBound int `json:"wlan_userage_upper_bound,omitempty"`

	// Only use for ads api
	AdsConfig string `json:"ads_config,omitempty"`
	AdsKeyword string `json:"ads_keyword,omitempty"`
	AdsSource string `json:"ads_source,omitempty"`
	AdsType string `json:"ads_type,omitempty"`
	Ads []map[string]interface{} `json:"ads,omitempty"`
	AdsBatchId string `json:"ads_batchid,omitempty"`
	AdsBatchSize string `json:"ads_batchsize,omitempty"`

	// Only use for conversion api
	AdsSeqNum string `json:"ads_seqnum,omitempty"`
	AdsClickUrl string `json:"ads_clickurl,omitempty"`
}
