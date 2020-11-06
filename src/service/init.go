/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-22
*/
package service

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"jjservice/src/middleware"

	"jjservice/src/logger"
	"jjservice/src/util"
)

func InitRouter() *gin.Engine {
	// 设置模式
	f, err := util.GetLog()
	if err != nil {
		gin.SetMode(gin.DebugMode)
		logger.JJLogger.Print("加载本地日志错误", err)
		gin.DefaultWriter = os.Stdout
		// 默认打印到console
	}else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		// 测试时打印到双端 生产环境打印到日志里
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	// 初始化视图
	var r *gin.Engine
	r = gin.New()
	// 中间件
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(util.CustomLog))
	r.Use(middleware.Cors())
	r.Use(middleware.ErrorWrapper())
	r.Use(middleware.ForbinddenWrapper())
	r.Use(middleware.AuthWrapper())

	// 添加路由
	initDefault(r)
	initApiStatus(r)
	initApiStart(r)
	initApiStop(r)
	initApiLog(r)
	initApiBackUp(r)
	return r
}

func InitDevRouter() *gin.Engine {
	// 设置模式
	f, err := util.GetLog()
	if err != nil {
		gin.SetMode(gin.DebugMode)
		fmt.Println("加载本地日志错误", err.Error())
		gin.DefaultWriter = os.Stdout
		// 默认打印到console
	}else {
		gin.SetMode(gin.DebugMode)
		gin.DisableConsoleColor()
		// 测试时打印到双端 生产环境打印到日志里
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	// 初始化视图
	var r *gin.Engine
	r = gin.New()
	// 中间件
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(util.CustomLog))
	r.Use(middleware.Cors())
	r.Use(middleware.ErrorWrapper())
	//r.Use(middleware.ForbinddenWrapper())
	//r.Use(middleware.AuthWrapper())

	// 添加路由
	initDefault(r)
	initApiStatus(r)
	initApiStart(r)
	initApiStop(r)
	initApiLog(r)

	return r
}