/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-26
*/
package util

import (
	"strconv"
	"strings"
)

func TokenDecrypt(tokenRaw string) (token string, timeStamp int64) {
	tmp := strings.Split(tokenRaw, "+")
	if len(tmp) > 1 {
		token = tmp[0]
		timeStamp, _ = strconv.ParseInt(tmp[1], 10, 64)
		return
	}
	return
}
