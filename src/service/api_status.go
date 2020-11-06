/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-23
*/

/*
状态路由模块
承载获取本地服务工作状态的功能
包括是否运行，运行PID，运行占用内存大小
注意pgrep和ps | grep区分 flask应用不能使用pgrep获取参数
*/
package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"jjservice/src/logger"
	"jjservice/src/model"
	"jjservice/src/model/shell"
	"jjservice/src/util"
)

// 状态路由
func initApiStatus(r *gin.Engine) {
	defaultHandler := func(c *gin.Context) {
		util.JJResponse(
			c,
			http.StatusOK,
			"this's API for Status",
			map[string]string{"api": "ApiStatus", "path": c.FullPath()},
		)
	}

	ApiStatus := r.Group("/api/status")
	{
		ApiStatus.GET("", defaultHandler)
		ApiStatus.POST("/all", getAllApps)
		ApiStatus.POST("/sys", getSys)
		ApiStatus.POST("/app/:app_id", getApp)
		ApiStatus.POST("/jjservice", statusJJService)
		ApiStatus.POST("/jjgo", statusJJGo)
		ApiStatus.POST("/jjbox", statusJJBox)
		ApiStatus.POST("/mysite", statusMysite)
		ApiStatus.POST("/blog", statusBlog)
		ApiStatus.POST("/jjmail", statusJJMail)
		ApiStatus.POST("/mgek", statusMgek)
		ApiStatus.POST("/mgekdoc", statusDoc)
		ApiStatus.POST("/mgekfile", statusMgekFile)
		ApiStatus.POST("/mgekcloud", statusNginx)
		ApiStatus.POST("/resume", statusNginx)
		ApiStatus.POST("/holding", statusNginx)
	}
}

// 获取全部服务信息，详情页为状态所以无需请求多个接口
func getAllApps(c *gin.Context) {
	util.JJResponse(
		c,
		http.StatusOK,
		"服务列表获取成功",
		util.ALL_APPS,
	)
}

// 获取系统信息
func getSys(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_SYS)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"刷新系统信息失败",
			model.FAILED,
		)
		return
	}
	data := strings.Split(opt,"|")
	util.JJResponse(
		c,
		http.StatusOK,
		"刷新系统信息成功",
		data,
	)
}
// 根据app_id获取app详细信息
func getApp(c *gin.Context) {
	appId := c.Param("app_id")
	app, err := util.ALL_APPS[appId]
	if !err {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务详情获取失败",
			app,
		)
		return
	}
	util.JJResponse(
		c,
		http.StatusOK,
		"服务详情获取成功",
		app,
	)
}

// 状态的判断逻辑很简单 只需要判断是否在运行即可
func statusJJService(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_JJService)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// jjgo的服务状态
func statusJJGo(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_JJGO)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// JJBox的状态
func statusJJBox(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_JJBOX)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// renj.io站点状态
func statusMysite(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_MYSITE)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// 博客状态
func statusBlog(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_BLOG)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// celery状态
func statusJJMail(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_JJMAIL)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// Mgek状态
func statusMgek(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_MGEK)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// MgekFile状态
func statusMgekFile(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_NGINX)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}

// 文档使用的是mgek接口
func statusDoc(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_MGEKDOC)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}


// Nginx通用状态
func statusNginx(c *gin.Context) {
	opt, err := util.RunShell(shell.STATUS_SERVICE_NGINX)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.ALIVE,
		)
		return
	}else{
		util.JJResponse(
			c,
			http.StatusOK,
			"服务状态获取成功",
			model.DEAD,
		)
		return
	}
}