/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-26
*/
package util

import (
	"strconv"
	"time"
)

func TokenEncrypt() string {
	timeNow := time.Now().Unix()
	timeNow2String := strconv.FormatInt(timeNow + Conf.TimeDiff, 10)
	token := Conf.CrsfToken + "+" + timeNow2String
	return token
}
