/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-21
*/
package model

import "time"

type Configs struct {
	// 应用配置
	AppId 			string
	Port   			int
	Mode            string
	ReadTimeout 	time.Duration
	WriteTimeout 	time.Duration
	// 日志配置
	LogPath 		string
	// 验证配置
	Referrer 		string
	CrsfToken 		string
	Host 			string
	TimeDiff        int64
}
