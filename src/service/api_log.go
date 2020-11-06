/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-22
*/

/*
日志路由模块
承载删除本地日志的功能，当前仅支持手动删除，调用echo "" >来清空日志
建议使用文件IO操作清除日志，避免cmd执行的阻塞
为了保证响应 不管有没有成功都返回值 所以应该使用非阻塞的cmd
读取使用get 删除使用post
 */

package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"jjservice/src/logger"
	"jjservice/src/middleware"
	"jjservice/src/model"
	"jjservice/src/model/shell"
	"jjservice/src/util"
)

// 日志路由
func initApiLog(r *gin.Engine) {
	defaultHandler := func(c *gin.Context) {
		util.JJResponse(
			c,
			http.StatusOK,
			"this's API for Log",
			map[string]string{"api": "ApiLog", "path": c.FullPath()},
			)
	}

	ApiLog := r.Group("/api/log")
	// 敏感路由需要验证
	ApiLog.Use(middleware.TokenWrapper())
	{
		ApiLog.GET("", defaultHandler)
		ApiLog.GET("/jjservice", getJJService)
		ApiLog.GET("/jjgo", getJJGo)
		ApiLog.GET("/jjbox", getJJBox)
		ApiLog.GET("/mysite", getMysite)
		ApiLog.GET("/blog", getBlog)
		ApiLog.GET("/jjmail", getJJMail)
		ApiLog.GET("/mgek", getMgek)

		// todo 删除日志应该遵循文件锁
		ApiLog.POST("/jjservice", delJJServiceLog)
		ApiLog.POST("/jjgo", delJJGo)
		ApiLog.POST("/jjbox", delJJBox)
		ApiLog.POST("/mysite", delMysite)
		ApiLog.POST("/blog", delBlog)
		ApiLog.POST("/jjmail", delJJMail)
		ApiLog.POST("/mgek", delMgek)
	}
}

// 删除jjservice服务的日志
func delJJServiceLog(c *gin.Context) {
	opt, err := util.RunShell(shell.DEL_SERVICE_JJService)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除成功",
			model.SUCCESS,
		)
		return
	}
}

// 删除jjgo服务的日志
func delJJGo(c *gin.Context) {
	opt, err := util.RunShell(shell.DEL_SERVICE_JJGO)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除成功",
			model.SUCCESS,
		)
		return
	}
}

// 删除jjbox服务的日志
func delJJBox(c *gin.Context) {
	opt, err := util.RunShell(shell.DEL_SERVICE_JJBOX)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除成功",
			model.SUCCESS,
		)
		return
	}
}

// 删除mysite服务的日志
func delMysite(c *gin.Context) {
	opt, err := util.RunShell(shell.DEL_SERVICE_MYSITE)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除成功",
			model.SUCCESS,
		)
		return
	}
}

// 删除blog服务的日志
func delBlog(c *gin.Context) {
	opt, err := util.RunShell(shell.DEL_SERVICE_BLOG)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除成功",
			model.SUCCESS,
		)
		return
	}
}

// 删除celery服务的日志
func delJJMail(c *gin.Context) {
	opt, err := util.RunShell(shell.DEL_SERVICE_JJMAIL)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除成功",
			model.SUCCESS,
		)
		return
	}
}

// 删除mgek服务的日志
func delMgek(c *gin.Context) {
	opt, err := util.RunShell(shell.DEL_SERVICE_MGEK)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志删除成功",
			model.SUCCESS,
		)
		return
	}
}

// GET func
func getJJService(c *gin.Context) {
	getLog(c, shell.LOG_SERVICE_JJService)
}

func getJJGo(c *gin.Context) {
	getLog(c, shell.LOG_SERVICE_JJGO)
}

func getJJBox(c *gin.Context) {
	getLog(c, shell.LOG_SERVICE_JJBOX)
}

func getMysite(c *gin.Context) {
	getLog(c, shell.LOG_SERVICE_MYSITE)
}

func getBlog(c *gin.Context) {
	getLog(c, shell.LOG_SERVICE_BLOG)
}

func getJJMail(c *gin.Context) {
	getLog(c, shell.LOG_SERVICE_JJMAIL)
}

func getMgek(c *gin.Context) {
	getLog(c, shell.LOG_SERVICE_MGEK)
}

// 尝试统一函数处理
func getLog(c *gin.Context, cmd string) {
	opt, err := util.RunShell(cmd)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"日志获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) <= 0 || opt == "" || strings.Contains(opt, "No such file or directory") {
		util.JJResponse(
			c,
			http.StatusOK,
			"日志获取失败",
			model.FAILED,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"日志获取成功",
			opt,
		)
		return
	}
}