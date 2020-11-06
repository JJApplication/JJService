/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-23
*/

/*
JJResponse Restful的接口
JJFile 二进制文件接口
JJHtml html模板渲染
 */
package util

import "github.com/gin-gonic/gin"

func JJResponse(ctx *gin.Context, httpCode int, msg string, data interface{}) {
	ctx.JSON(
			httpCode,
			gin.H{
				"msg": msg,
				"data": data,
			},
		)
}

func JJFile(ctx *gin.Context, filepath string)  {
	ctx.File(filepath)
}

func JJHtml(ctx *gin.Context, html string)  {
	ctx.HTML(200, html, "")
}