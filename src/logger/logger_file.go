/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-02
*/
package logger

import (
	"os"
	"jjservice/src/util"
)

// 读取本地日志文件流
func getLog() (*os.File, error) {
	logPath := util.Conf.LogPath
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