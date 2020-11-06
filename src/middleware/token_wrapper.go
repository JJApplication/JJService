/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-26
*/
package middleware

import (
	"net/http"
	"time"

	"jjservice/src/util"

	"github.com/gin-gonic/gin"
)

// token信息的解密，token包含tokenstring + UINX时间戳
func TokenWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRaw := c.Request.Header.Get("token")
		if tokenRaw != "" {
			token, timeStamp := util.TokenDecrypt(tokenRaw)
			timeNow := time.Now().Unix()
			timeDiff := timeNow - timeStamp
			// 差值不应该大于配置

			if token == util.Conf.CrsfToken && timeDiff <= util.Conf.TimeDiff {
				c.Next()
			}else if token == util.Conf.CrsfToken && timeDiff > util.Conf.TimeDiff {
				util.JJResponse(
					c,
					http.StatusUnauthorized,
					// 因为想要输出错误页面信息
					"token expired",
					"authorization failed",
				)
				c.Abort()
				return
			}else {
				util.JJResponse(
					c,
					http.StatusUnauthorized,
					// 因为想要输出错误页面信息
					"you don't have permission!",
					"authorization failed",
				)
				c.Abort()
				return
			}
		}else {
			util.JJResponse(
				c,
				http.StatusUnauthorized,
				// 因为想要输出错误页面信息
				"please add token with request",
				"authorization failed",
			)
			c.Abort()
			return
		}
	}
}

