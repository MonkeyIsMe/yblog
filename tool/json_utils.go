package tool

import (
	"encoding/json"
	"log"
	"yblog/constant"
)

// UnmarshalPageCookie 解析pageCookie
func UnmarshalPageCookie(cookie string) (page *constant.PageCookie, err error) {
	page = &constant.PageCookie{}
	if len(cookie) > 0 {
		// 无效cookie
		if !json.Valid([]byte(cookie)) {
			log.Fatalf("cookie不合法")
			return
		}

		err = json.Unmarshal([]byte(cookie), page)
		if err != nil {
			log.Fatalf("PageCookie Unmarshal err: [%+v]", err)
			return
		}
	}
	return
}
