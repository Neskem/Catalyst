# catalyst API(v1)
## Usage Verification Rule

- detect_bot_by_http_user_agent -> IsBot Exception
  user agent list: ["Googlebot", "Baiduspider", "bingbot", "msnbot", "YandexMobileBot", "PhantomJS", "web spider", "Sogou web spider", "AppEngine-Google", "360Spider", "js bot"]
- detect_bot_by_requests_count_from_ip+endpoint+http_user_agent -> IsOverRateLimit Exception
  user agent ignorelist: ["GoogleStackdriverMonitoring", "GoogleHC", "BT_DATA_SERVICE_TEAM_STRESS_TEST_ONLY"]

## page_id Cal Rule

- 處理邏輯:
  1.  decode url_pageid
  2.  移除 url 上的 # 之後的資料
  3.  移除 url 上的 scheme(http://, https://)
  4.  將 3. 的結果 做 SHA1
- ex:
  1.  url_pageid: https%3A%2F%2Fzi.media%2F%40abrabbit%2Fpost%2FGrvKPz
  2.  url_pageid(decode): https://zi.media/@abrabbit/post/GrvKPz
  3.  page_id: 79ab6662597db63ddfddc92697190d7e744da7f3

## Get Viewer(track id/partner cookie)

> - **Request**
>    Endpoint: `GET /v1/cat_trid`
>    Content-Type: `image/png (application/json)`
> - **Response**
>   - cat_trid: catalyst cookie 內容, 如果不存在則產生新的. str(uuid.uuid4())+"."+str(time.time())
>   - pt_jwt: 夥伴系統 cookie(pt_jwt)
>   ```json
>   {
>       "cat_trid": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       "pt_jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJicmVha3RpbWUiLCJpYXQiOjE1MzA2ODE4MzQsInVzZXJfbmFtZSI6IkZYOFYyMTgiLCJkb21haW5zIjpbInN0eWxpc2htYWtlLmNvbSIsImJsb2cudWRuLmNvbS85Zjk4ZmE1OSJdLCJleHAiOjE1MzA3NjgyMzQsInBhcnRuZXJJZCI6IkZYOFYyMTgifQ.jOXYKvCs0IDoF4yxv5tcusy7SxDTlgUWC1Fza9VbPPs"
>   };
>   ```
> - **Example**
>   ```bash
>   curl -v -X GET \
>   http://stg-catalyst.breaktime.com.tw/v1/cat_trid \
>   -H 'cache-control: no-cache' \
>   -H 'content-type: application/json'
>   ```

## Collection Viewer(footprint)

> - **Request**
>    Endpoint: `POST /v1/footprint`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "txn_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>       "spj": "ykw,???,???",
>       "is_infinity": 1,
>       "infinity_batchid": "b_1",/* firstpage_txn_id2 */
>       "infinite_url_firstpage": "http://test-web.zi.media/@damon624/post/Wg7wfY",/* firstpage_url */
>       "infinite_url": "http://test-web.zi.media/@damon624pixnetnetblog/post/123456",/* current_url */
>       "infinity_content_seqid": 2,
>       "infinity_content_percentage": 100
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
>   - 目前會有 ["http://... ... ..."
>     , "https://... ... ..."
>     , "applewebdata://003C13A2-7D1B-4284-9D4C-8725CB279E79"
>     , "content://0@media/external/file/10301"
>     , "content://com.android.providers.downloads.documents/document/151"
>     , "content://downloads/all_downloads/4681"
>     , "file://localhost/Users/tangomycin/Desktop/大阪/【大阪景點】Kids Plaza Osaka （大阪兒童樂園）：大阪親子遊必訪，超好玩的兒童天堂。 - 👒Mimi 韓の愛旅生活 👒.htm"
>     , "unmht://unmht/file.5/A:/Hi-End/★ [TAS]/TAS List - 古典黑膠共和國 ( ClassicalVinylRepublic ) - 樂多日誌.mht/"]
> - **Response**: None
> - **Example**
>
>   ```bash
>   curl -v -X POST \
>   http://stg-catalyst.breaktime.com.tw/v1/footprint \
>   -H 'cache-control: no-cache' \
>   -H 'content-type: application/json' \
>   -d '{
>   "fp": "d712112f431f62a13d12e201dcef0c37",
>   "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>   "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>   "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>   "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>   "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>   "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>   "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>   "referrer": "https://ck101.com/",
>   "device_type": "XiaoMi HM NOTE 1W",
>   "spj":"ykw,???,???"
>   }'
>   ```
>
> - **Redis(footprint)**
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>       "spj":"ykw,???,???",
>
>       "ip": "60.248.166.160",
>       "page_id": "4ecdd928256e4992d11ac6cba80acc875e071fdd",
>       "creation_time": "2018-07-06T10:12:02 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid_scheme": "http",
>       "url_pageid_hostname": "m.comic.ck101.com",
>       "url_pageid_port": None,
>       "url_pageid_path": "/vols/28687065/22",
>       "url_pageid_query": "'v1=A&V2=B'",
>       "url_pageid_query_args": {
>           "v1": "A",
>           "v2": "B",
>       },
>       "url_pageid_fragment": "2013_m_comic_allpage_320x50_top",
>
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_scheme": "http",
>       "url_hostname": "m.comic.ck101.com",
>       "url_port": None,
>       "url_path": "/vols/28687065/22",
>       "url_query": "",
>       "url_query_args": {},
>       "url_fragment": "2013_m_comic_allpage_320x50_top",
>
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical_scheme": "http",
>       "url_canonical_hostname": "m.comic.ck101.com",
>       "url_canonical_port": None,
>       "url_canonical_path": "/vols/28687065/22",
>       "url_canonical_query": "",
>       "url_canonical_query_args": {},
>       "url_canonical_fragment": "2013_m_comic_allpage_320x50_top",
>
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme": "http",
>       "url_og_hostname": "m.comic.ck101.com",
>       "url_og_port": None,
>       "url_og_path": "/vols/28687065/22",
>       "url_og_query": "",
>       "url_og_query_args": {},
>       "url_og_fragment": "2013_m_comic_allpage_320x50_top",
>
>       "referrer(decode)": "http://ck101.com/",
>       "referrer_scheme": "http",
>       "referrer_hostname": "ck101.com",
>       "referrer_port": None,
>       "referrer_path": "/",
>       "referrer_query": "",
>       "referrer_query_args": {},
>       "referrer_fragment": "",
>   ```
>
>
>       "ua":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36",
>       "ua_browser_family": "Chrome",
>       "ua_browser_version_major": 67,
>       "ua_browser_version_minor": 0,
>       "ua_browser_version_build": 3396,
>       "ua_browser_version_string": "67.0.3396.87",
>
>       "ua_os_family": "Mac OS X",
>       "ua_os_version_major": 10,
>       "ua_os_version_minor": 13,
>       "ua_os_version_build": 5,
>       "ua_os_version_string": "10.13.5",
>
>       "ua_device_family": "Other",
>       "ua_device_brand": None,
>       "ua_device_model": None,
>       "ua_is_mobile": False,
>       "ua_is_tablet": False,
>       "ua_is_touch_capable": False,
>       "ua_is_pc": True,
>       "ua_is_bot": False,
>
>       "ip": "127.0.0.1",
>       "ip_X-Forwarded-For": "",
>       "ip_X-Real-Ip": ""
>
> };
>
> ```
>
> ```

## Collection Viewer(yahoo age/gender)

> - **Request**
>    Endpoint: `POST /v1/profile_oath`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "user_age": 25,
>       "user_gender": "f",
>       "user_country": "tw"
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(profile_oath)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "user_age": 25,
>       "user_gender": "F",
>       "user_country": "TW",
>       "user_age(int)": 25,
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567"
>   };
>   ```

## Collection Viewer(fb/line id)

> - **Request**
>    Endpoint: `POST /v1/open_id`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "uid_line": "",
>       "uid_fb": ""
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(open_id)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "uid_line": "",
>       "uid_fb": "",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567"
>   };
>   ```

## Collection Viewer(wifi)

> - **Request**
>    Endpoint: `POST /v1/wifi`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "wlan_server": "cht8.cool3c.com",
>       "wlan_ip": "10.19.53.5",
>       "wlan_mac": "e4:8d:8c:08:de:96",
>       "wlan_gwaddr": "10.16.0.1:2060",
>       "wlan_token": "0e3f532d0d0db689def785a12ebbf979",
>       "wlan_timekey": "5810a114ddd211a747c1ca7ac31cb03c",
>       "wlan_userage": "???-???",
>       "wlan_usergender": "",
>       "wlan_usercategory": "分類1,分類2,分類3",
>       "wlan_access_loc": "玉泉里民活動中心",
>       "wlan_access_loc_region": "展演場所,台北市大同區,義_台北火車站_2KM_0625,義_台北西門町商圈_2KM_0625,西門町",
>       "wlan_device_loc": "台北市大同區環河北路一段96號1樓"
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(wifi)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "wlan_server": "cht8.cool3c.com",
>       "wlan_ip": "10.19.53.5",
>       "wlan_mac": "e4:8d:8c:08:de:96",
>       "wlan_gwaddr": "10.16.0.1:2060",
>       "wlan_token": "0e3f532d0d0db689def785a12ebbf979",
>       "wlan_timekey": "5810a114ddd211a747c1ca7ac31cb03c",
>       "wlan_userage": "",
>       "wlan_usergender": "",
>       "wlan_usercategory": "分類1,分類2,分類3",
>       "wlan_access_loc": "玉泉里民活動中心",
>       "wlan_access_loc_region": "展演場所,台北市大同區,義_台北火車站_2KM_0625,義_台北西門町商圈_2KM_0625,西門町",
>       "wlan_device_loc": "台北市大同區環河北路一段96號1樓",
>       "wlan_userage_lower_bound": "???",
>       "wlan_userage_upper_bound": "???",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567"
>   };
>   ```

## Collection Viewer(ads)

> - **Request**
>    Endpoint: `POST /v1/ads`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "ads_config":"",
>       "ads_keyword":"吊燈",
>       "ads_source":"ypa",
>       "ads_type":"",
>       "ads": [
>           {
>               "clickURL":"https://1r.search.yahoo.com/cbclk/dWU9Q0Q1MzBDQ0Q0MzZGNDVEMCZ1dD0xNTMwMjUyMDg4MDMwJnVvPTc0ODM1NTE2ODA3NTAzJmx0PTImZXM9WUNMSU56TUdQUzlRc0dtOA--/RV=2/RE=1530280888/RO=10/RU=https://www.bing.com/aclick?ld=d3VidoiAKrzPvrrSuvY0jz5jVUCUzHlWtnBPue1P7UxbzdS8mj68EMG_mjRNeykhm5qr0R1JKH9-YKKbkoAtGZAqB2e12lv8Wzqjnt4wn0z0aXS-lI4Wr6NeYoKDscri7iQFsEX08-QLuOZKuDkmGU3Z9kFBo&u=http%3a%2f%2fwww.cheaplight.com.tw%2fhtml%2fprodclass.php%3fprodClassNo%3d2/RK=2/RS=4d0KPKg.V1PHu.F3jwB7hSUUV0M-",
>               "descr":"1一級水晶球，出貨比展示品更新更亮，吊燈品質童叟無欺",
>               "sitehost":"1www.cheaplight.com.tw",
>               "h_img":"1#",
>               "title":"1難以置信的快樂價格-吊燈 - 高級吊燈保證便宜買到"
>           },
>           {
>               "clickURL":"https://1r.search.yahoo.com/cbclk/dWU9Q0Q1MzBDQ0Q0MzZGNDVEMCZ1dD0xNTMwMjUyMDg4MDMwJnVvPTc0ODM1NTE2ODA3NTAzJmx0PTImZXM9WUNMSU56TUdQUzlRc0dtOA--/RV=2/RE=1530280888/RO=10/RU=https://www.bing.com/aclick?ld=d3VidoiAKrzPvrrSuvY0jz5jVUCUzHlWtnBPue1P7UxbzdS8mj68EMG_mjRNeykhm5qr0R1JKH9-YKKbkoAtGZAqB2e12lv8Wzqjnt4wn0z0aXS-lI4Wr6NeYoKDscri7iQFsEX08-QLuOZKuDkmGU3Z9kFBo&u=http%3a%2f%2fwww.cheaplight.com.tw%2fhtml%2fprodclass.php%3fprodClassNo%3d2/RK=2/RS=4d0KPKg.V1PHu.F3jwB7hSUUV0M-",
>               "descr":"1一級水晶球，出貨比展示品更新更亮，吊燈品質童叟無欺",
>               "sitehost":"1www.cheaplight.com.tw",
>               "h_img":"1#",
>               "title":"1難以置信的快樂價格-吊燈 - 高級吊燈保證便宜買到"
>           }
>       ]
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(ads)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "ads_config":"",
>       "ads_keyword":"吊燈",
>       "ads_source":"ypa",
>       "ads_type":"",
>
>       "ads_batchid": "'9b12ed96-f350-4714-abdd-0371205464c1'",
>       "ads_batchsize": "2",
>       "ads_clickURL": "'https://1r.search.yahoo.com/cbclk/dWU9Q0Q1MzBDQ0Q0MzZGNDVEMCZ1dD0xNTMwMjUyMDg4MDMwJnVvPTc0ODM1NTE2ODA3NTAzJmx0PTImZXM9WUNMSU56TUdQUzlRc0dtOA--/RV=2/RE=1530280888/RO=10/RU=https://www.bing.com/aclick?ld=d3VidoiAKrzPvrrSuvY0jz5jVUCUzHlWtnBPue1P7UxbzdS8mj68EMG_mjRNeykhm5qr0R1JKH9-YKKbkoAtGZAqB2e12lv8Wzqjnt4wn0z0aXS-lI4Wr6NeYoKDscri7iQFsEX08-QLuOZKuDkmGU3Z9kFBo&u=http%3a%2f%2fwww.cheaplight.com.tw%2fhtml%2fprodclass.php%3fprodClassNo%3d2/RK=2/RS=4d0KPKg.V1PHu.F3jwB7hSUUV0M-'",
>       "ads_descr":"'1一級水晶球，出貨比展示品更新更亮，吊燈品質童叟無欺'",
>       "ads_sitehost":"'1www.cheaplight.com.tw'",
>       "ads_h_img":"1#",
>       "ads_title":"'1難以置信的快樂價格-吊燈 - 高級吊燈保證便宜買到'",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567",
>   };
>   ```

## Collection Viewer(conversion)

> - **Request**
>    Endpoint: `POST /v1/conversion`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "ads_seqnum": "1",
>       "ads_clickurl": "'https://1r.search.yahoo.com/cbclk/dWU9Q0Q1MzBDQ0Q0MzZGNDVEMCZ1dD0xNTMwMjUyMDg4MDMwJnVvPTc0ODM1NTE2ODA3NTAzJmx0PTImZXM9WUNMSU56TUdQUzlRc0dtOA--/RV=2/RE=1530280888/RO=10/RU=https://www.bing.com/aclick?ld=d3VidoiAKrzPvrrSuvY0jz5jVUCUzHlWtnBPue1P7UxbzdS8mj68EMG_mjRNeykhm5qr0R1JKH9-YKKbkoAtGZAqB2e12lv8Wzqjnt4wn0z0aXS-lI4Wr6NeYoKDscri7iQFsEX08-QLuOZKuDkmGU3Z9kFBo&u=http%3a%2f%2fwww.cheaplight.com.tw%2fhtml%2fprodclass.php%3fprodClassNo%3d2/RK=2/RS=4d0KPKg.V1PHu.F3jwB7hSUUV0M-'",
>       "ads_descr":"'1一級水晶球，出貨比展示品更新更亮，吊燈品質童叟無欺'",
>       "ads_sitehost":"'1www.cheaplight.com.tw'",
>       "ads_h_img":"1#",
>       "ads_title":"'1難以置信的快樂價格-吊燈 - 高級吊燈保證便宜買到'",
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(ads_seqnum, ads_clickURL), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(conversion)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>
>       "ads_seqnum": "1",
>       "ads_clickurl": "'https://1r.search.yahoo.com/cbclk/dWU9Q0Q1MzBDQ0Q0MzZGNDVEMCZ1dD0xNTMwMjUyMDg4MDMwJnVvPTc0ODM1NTE2ODA3NTAzJmx0PTImZXM9WUNMSU56TUdQUzlRc0dtOA--/RV=2/RE=1530280888/RO=10/RU=https://www.bing.com/aclick?ld=d3VidoiAKrzPvrrSuvY0jz5jVUCUzHlWtnBPue1P7UxbzdS8mj68EMG_mjRNeykhm5qr0R1JKH9-YKKbkoAtGZAqB2e12lv8Wzqjnt4wn0z0aXS-lI4Wr6NeYoKDscri7iQFsEX08-QLuOZKuDkmGU3Z9kFBo&u=http%3a%2f%2fwww.cheaplight.com.tw%2fhtml%2fprodclass.php%3fprodClassNo%3d2/RK=2/RS=4d0KPKg.V1PHu.F3jwB7hSUUV0M-'",
>       "ads_descr":"'1一級水晶球，出貨比展示品更新更亮，吊燈品質童叟無欺'",
>       "ads_sitehost":"'1www.cheaplight.com.tw'",
>       "ads_h_img":"1#",
>       "ads_title":"'1難以置信的快樂價格-吊燈 - 高級吊燈保證便宜買到'",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567",
>   };
>   ```

## Collection Viewer Highlighted Text

> - **Request**
>    Endpoint: `POST /v1/highlighted_text`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "highlightedtext": ""
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
>   - 必填欄位(至少一項) ["highlightedtext"] -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(highlighted_text)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "highlightedtext": "",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567"
>   };
>   ```

## Collection Viewer Openlink

> - **Request**
>    Endpoint: `POST /v1/openlink`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "openlink_href": "",
>       "openlink_text": "",
>       "openlink_img_alt": "",
>       "openlink_img_src": ""
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
>   - 必填欄位(至少一項) ["openlink_href", "openlink_text", "openlink_img_alt", "openlink_img_src"] -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(openlink)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "openlink_href": "",
>       "openlink_text": "",
>       "openlink_img_alt": "",
>       "openlink_img_src": "",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567"
>   };
>   ```

## Collection Viewer Session Stay

> - **Request**
>    Endpoint: `POST /v1/session_stay`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "session_stay": 60
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(fp, txn_id, url_pageid), 且 url_pageid 欄位值需可正常 parsing -> MissingRequiredFields Exception
>   - 必填欄位(至少一項) ["session_stay"] -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(session_stay)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "session_stay": 60,
>       "session_stay(int)": 60,
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567"
>   };
>   ```

## Collection Viewer(javascript error)

> - **Request**
>    Endpoint: `POST /v1/js_err`
>    Content-Type: `image/png (application/json)`
>
>   ```json
>   {
>       "fp": "d712112f431f62a13d12e201dcef0c37",
>       "txn_id": "d712112f431f62a13d12e201dcef0c37???",
>
>       /* session_id 內容應為 /v1/cat_trid 取得的 cat_trid, 如果使用者清空 cookie, 會欄位會是新的 session_id */
>       "session_id": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045",
>       /* txn_id 內容應為 cat_trid 加上 ???, 超過 30 分即使在同一頁面, ??? 也會更新 */
>       "txn_id2": "a0ab7276-3710-4b26-822e-5a0ac42485f0.1530678303.1034045_???",
>
>       "url": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og" : "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_pageid": "http://m.comic.ck101.com/vols/28687065/22?v1=A&V2=B#2013_m_comic_allpage_320x50_top",
>       "referrer": "https://ck101.com/",
>       "device_type": "XiaoMi HM NOTE 1W",
>
>       "err_msg": ""
>   };
>   ```
>
> - **Request Validation**
>   - 必填欄位(至少一項) ["err_msg"] -> MissingRequiredFields Exception
> - **Response**: None
> - **Redis(js_err)**
>
>   ```json
>   {
>       ... ... .../* 同 request */
>       "err_msg": "",
>
>       "url_pageid(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_canonical(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "url_og_scheme(decode)": "http://m.comic.ck101.com/vols/28687065/22#2013_m_comic_allpage_320x50_top",
>       "referrer(decode)": "http://ck101.com/",
>       "ip": "127.0.0.1",
>       "page_id": "0069d7b53ab01318ab063774e8af584ca6cb1767",
>       "creation_time": "2018-07-09T14:16:46 +0800",
>       "hbase_rowkey": "a_2018-07-06T10:12:02.1234567"
>   };
>   ```
