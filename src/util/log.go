/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-23
*/
package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// 读取本地日志文件流
func GetLog() (*os.File, error) {
	conf, _ := ReadConf()
	logPath := conf.LogPath
	// 先判断是否存在
	// 不用判断
	_ ,isExist := os.Stat(logPath)
	if isExist != nil {
		// 不存在
		file, _ := os.Create(logPath)
		return file, nil
	}else{
		file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}

		return file, nil
	}
}

func CustomLog(param gin.LogFormatterParams) string{
	return fmt.Sprintf("[JJService] %s - [%s] \" %s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.ErrorMessage,
	)
}

func Logger(c *gin.Context, err error)  {
	f, _ := GetLog()
	if err != nil {
		fmt.Fprintf(f, "[ERROR]@API{%s} %s\n",c.FullPath(), err.Error())
	}
}