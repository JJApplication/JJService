/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-24
*/

/*
启动应用路由模块
承载启动本地app功能
使用shell脚本形式启动本地程序
需要注意start前保证此服务没有运行否则不执行操作
*/
package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"jjservice/src/middleware"
	"jjservice/src/model"
	"jjservice/src/model/shell"
	"jjservice/src/util"
)

// 启动路由
func initApiStart(r *gin.Engine) {
	defaultHandler := func(c *gin.Context) {
		util.JJResponse(
			c,
			http.StatusOK,
			"this's API for Start Apps",
			map[string]string{"api": "ApiStart", "path": c.FullPath()},
		)
	}

	ApiStart := r.Group("/api/start")
	ApiStart.Use(middleware.TokenWrapper())
	{
		ApiStart.GET("", defaultHandler)
		ApiStart.POST("/jjservice", startJJService)
		ApiStart.POST("/jjgo", startJJGo)
		ApiStart.POST("/jjbox", startJJBox)
		ApiStart.POST("/mysite", startMysite)
		ApiStart.POST("/blog", startBlog)
		ApiStart.POST("/jjmail", startJJMail)
		ApiStart.POST("/mgek", startMgek)
	}
}

// 开启jjservice 注意在jjservice被关闭后此服务是无法运行的
// 执行结果为空说明没有打印nohup.out
func startJJService(c *gin.Context) {
	opt := util.RunShellAsync(shell.START_SERVICE_JJService)
	if !checkOpt(opt, shell.START_SERVICE_JJService) {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动成功",
			model.SUCCESS,
		)
		return
	}
}

// 开启jjgo
func startJJGo(c *gin.Context) {
	opt := util.RunShellAsync(shell.START_SERVICE_JJGO)
	if !checkOpt(opt, shell.START_SERVICE_JJGO) {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动成功",
			model.SUCCESS,
		)
		return
	}
}

// 开启jjbox
func startJJBox(c *gin.Context) {
	opt := util.RunShellAsync(shell.START_SERVICE_JJBOX)
	if !checkOpt(opt, shell.START_SERVICE_JJBOX) {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动成功",
			model.SUCCESS,
		)
		return
	}
}

// 开启mysite
func startMysite(c *gin.Context) {
	opt := util.RunShellAsync(shell.START_SERVICE_MYSITE)
	if !checkOpt(opt, shell.START_SERVICE_MYSITE) {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动成功",
			model.SUCCESS,
		)
		return
	}
}

// 开启mgek
func startMgek(c *gin.Context) {
	opt := util.RunShellAsync(shell.START_SERVICE_MGEK)
	if !checkOpt(opt, shell.START_SERVICE_MGEK) {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动成功",
			model.SUCCESS,
		)
		return
	}
}

// 开启blog
func startBlog(c *gin.Context) {
	opt := util.RunShellAsync(shell.START_SERVICE_BLOG)
	if !checkOpt(opt, shell.START_SERVICE_BLOG) {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动成功",
			model.SUCCESS,
		)
		return
	}
}

// 开启celery
func startJJMail(c *gin.Context) {
	opt := util.RunShellAsync(shell.START_SERVICE_JJMAIL)
	if !checkOpt(opt, shell.START_SERVICE_JJMAIL) {
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"应用启动成功",
			model.SUCCESS,
		)
		return
	}
}

func checkOpt(opt error, cmd string) bool {
	// 启动命令的opt检查
	if opt == nil {
		// 正常情况下不会为空
		return true
	}else {
		o := opt.Error()
		if strings.Contains(o, "status 1") || strings.Contains(o, "1"){
			// 状态码1 表示启动完成
			// 启动完成 关闭其bash进程
			_ =util.RunShellAsync(shell.STOP_BASH + " " + cmd)
			return true
		}else {
			// 状态码2 表示已经在运行
			return false
		}
	}
}