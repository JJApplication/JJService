/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-06
*/
package main
// 拥有热启动的jjservice

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"jjservice/src/logger"
	"jjservice/src/service"
	"jjservice/src/util"
)

func hotfix()  {
	info, _ := util.ReadAppInfo()

	logger.JJLogger.Print(fmt.Sprintf("service id %s", util.Conf.AppId), nil)
	logger.JJLogger.Print(fmt.Sprintf("service running at port: %v", util.Conf.Port), nil)
	logger.JJLogger.Print(fmt.Sprintf("service infos: @jjservice{By Landers}" +
		"app_id: %s\n" +
		"app_des: %s\n" +
		"app_version: %s\n" +
		"build_version: %s\n",
		info.App_id, info.App_des, info.App_version, info.App_build), nil)

	// 确定运行模式
	var r *gin.Engine
	switch util.Conf.Mode {
	case "production":
		r = service.InitRouter()
	case "debug":
		r = service.InitDevRouter()
	default:
		r = service.InitRouter()
	}

	s := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d",util.Conf.Port),
		Handler: r,
		ReadTimeout: util.Conf.ReadTimeout,
		WriteTimeout: util.Conf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// pipe
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			logger.JJLogger.Print("service going wrong %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	logger.JJLogger.Print("Shutdown Server ...", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.JJLogger.Print("Server Shutdown:", err)
	}

	logger.JJLogger.Print("Exit Server", nil)
}

