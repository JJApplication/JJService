/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-24
*/

/*
关闭应用路由模块
承载关闭本地服务功能
使用执行shell脚本的方式关闭，理想状态时使用PID的方式关闭
*/
package service

import (
	"github.com/gin-gonic/gin"
	"jjservice/src/middleware"
	"jjservice/src/model"
	"jjservice/src/model/shell"
	"jjservice/src/util"
	"net/http"
	"strings"
)

// 关闭路由
func initApiStop(r *gin.Engine) {
	defaultHandler := func(c *gin.Context) {
		util.JJResponse(
			c,
			http.StatusOK,
			"this's API for Stop Apps",
			map[string]string{"api": "ApiStop", "path": c.FullPath()},
		)
	}

	ApiStop:= r.Group("/api/stop")
	ApiStop.Use(middleware.TokenWrapper())
	{
		ApiStop.GET("", defaultHandler)
		ApiStop.POST("/jjservice", stopJJService)
		ApiStop.POST("/jjgo", stopJJGo)
		ApiStop.POST("/jjbox", stopJJBox)
		ApiStop.POST("/mysite", stopMysite)
		ApiStop.POST("/blog", stopBlog)
		ApiStop.POST("/jjmail", stopJJMail)
		ApiStop.POST("/mgek", stopMgek)
	}
}

func stopJJService(c *gin.Context) {
	stop(c, shell.STOP_SERVICE_JJService)
}

func stopJJGo(c *gin.Context) {
	stop(c, shell.STOP_SERVICE_JJGO)
}

func stopJJBox(c *gin.Context) {
	stop(c, shell.STOP_SERVICE_JJBOX)
}

func stopBlog(c *gin.Context) {
	stop(c, shell.STOP_SERVICE_BLOG)
}

func stopMysite(c *gin.Context) {
	stop(c, shell.STOP_SERVICE_MYSITE)
}

func stopMgek(c *gin.Context) {
	stop(c, shell.STOP_SERVICE_MGEK)
}

func stopJJMail(c *gin.Context) {
	stop(c, shell.STOP_SERVICE_JJMAIL)
}

func stop(c *gin.Context, cmd string) {
	opt, err := util.RunShell(cmd)
	if err != nil {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用停止失败",
			model.FAILED,
		)
		return
	}else {
		if strings.Contains(opt, "successfully stopped") {
			util.JJResponse(
				c,
				http.StatusOK,
				"应用停止成功",
				model.SUCCESS,
			)
			return
		}else {
			util.JJResponse(
				c,
				http.StatusOK,
				"应用停止失败",
				model.FAILED,
			)
			return
		}
	}
}