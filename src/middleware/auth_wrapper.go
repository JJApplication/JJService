/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-25
*/
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"jjservice/src/util"
)

// 一般验证中间件 提供refer host验证
func AuthWrapper() gin.HandlerFunc {
	conf := util.Conf
	return func(c *gin.Context) {
		// 优先判断refer
		ctxRefer := c.Request.Referer()
		// 替换全部http:// https://
		ctxRefer = strings.Replace(ctxRefer, "http://", "", -1)
		ctxRefer = strings.Replace(ctxRefer, "https://", "", -1)
		allowedRefer := strings.Fields(conf.Referrer)
		var flag bool = true
		for _, v := range allowedRefer {
			if strings.Contains(ctxRefer, v) {
				flag = false
			}
		}
		if flag {
			util.JJResponse(
				c,
				http.StatusUnauthorized,
				// 因为想要输出错误页面信息
				"you don't have permission!",
				"referrers is not allowed",
			)
			c.Abort()
			return
		}
		// 判断Host
		ctxHost := c.Request.Host
		if !strings.Contains(conf.Host, ctxHost) {
			util.JJResponse(
				c,
				http.StatusUnauthorized,
				// 因为想要输出错误页面信息
				"you don't have permission!",
				"hosts is not allowed",
			)
			c.Abort()
			return
		}
		c.Next()
	}
}
