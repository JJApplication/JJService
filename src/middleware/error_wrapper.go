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

// 默认的404处理中间件
func ErrorWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errCode := c.Writer.Status()
		if errCode == 404 {
			util.JJResponse(
				c,
				http.StatusNotFound,
				// 因为想要输出错误页面信息
				"can't find what you are looking for!",
				"404 Not Found",
				)
		}
	}
}
