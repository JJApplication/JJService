/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-24
*/
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jjservice/src/util"
)
// 默认的403响应中间件，并不是验证中间件
func ForbinddenWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errCode := c.Writer.Status()
		if errCode == 403 {
			util.JJResponse(
				c,
				http.StatusForbidden,
				// 因为想要输出错误页面信息
				"you don't have permission!",
				"403 Forbidden",
			)
		}
	}
}
