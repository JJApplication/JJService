/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-04
*/
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jjservice/src/logger"
	"jjservice/src/middleware"
	"jjservice/src/model"
	"jjservice/src/model/shell"
	"jjservice/src/util"
)

// 用于备份二进制文件的脚本
// 对于python的动态脚本使用压缩的方式备份
func initApiBackUp(r *gin.Engine) {
	defaultHandler := func(c *gin.Context) {
		util.JJResponse(
			c,
			http.StatusOK,
			"this's API for BackUp",
			map[string]string{"api": "ApiBackUp", "path": c.FullPath()},
		)
	}
	ApiBackUp := r.Group("/api/backup")
	// 敏感路由需要验证
	ApiBackUp.Use(middleware.TokenWrapper())
	{
		ApiBackUp.GET("", defaultHandler)
		ApiBackUp.POST("/jjservice", backupJJService)
		ApiBackUp.POST("/jjgo", backupJJGo)
		ApiBackUp.POST("/jjbox", backupJJBox)
		ApiBackUp.POST("/mysite", backupMysite)
		ApiBackUp.POST("/blog", backupBlog)
		ApiBackUp.POST("/jjmail", backupJJMail)
		ApiBackUp.POST("/mgek", backupMgek)
	}

}

// 备份jjservice
func backupJJService (c *gin.Context) {
	opt, err := util.RunShell(shell.BACKUP_SERVICE_JJService)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
	}else {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份成功",
			model.SUCCESS,
		)
	}
}

// 备份jjgo
func backupJJGo (c *gin.Context) {
	opt, err := util.RunShell(shell.BACKUP_SERVICE_JJGO)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
	}else {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份成功",
			model.SUCCESS,
		)
	}
}

// 备份jjbox
func backupJJBox (c *gin.Context) {
	opt, err := util.RunShell(shell.BACKUP_SERVICE_JJBOX)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
	}else {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份成功",
			model.SUCCESS,
		)
	}
}

// 备份mysite
func backupMysite (c *gin.Context) {
	opt, err := util.RunShell(shell.BACKUP_SERVICE_MYSITE)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
	}else {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份成功",
			model.SUCCESS,
		)
	}
}

// 备份blog
func backupBlog (c *gin.Context) {
	opt, err := util.RunShell(shell.BACKUP_SERVICE_BLOG)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
	}else {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份成功",
			model.SUCCESS,
		)
	}
}

// 备份celery
func backupJJMail (c *gin.Context) {
	opt, err := util.RunShell(shell.BACKUP_SERVICE_JJMAIL)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
	}else {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份成功",
			model.SUCCESS,
		)
	}
}

// 备份mgek
func backupMgek(c *gin.Context) {
	opt, err := util.RunShell(shell.BACKUP_SERVICE_MGEK)
	if err != nil {
		logger.JJLogger.Context(c, err)
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
		return
	}
	if len(opt) >= 1 {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份失败",
			model.FAILED,
		)
	}else {
		util.JJResponse(
			c,
			http.StatusOK,
			"备份成功",
			model.SUCCESS,
		)
	}
}