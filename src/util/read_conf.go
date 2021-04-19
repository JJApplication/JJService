/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-21
*/
package util

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-ini/ini"
	"jjservice/src/model"
)

var Conf model.Configs

func init() {
	Conf, _ = ReadConf()
}

func ReadConf() (model.Configs, error) {
	var conf model.Configs
	cwdPath , pathErr := os.Getwd()
	if pathErr != nil {
		return model.Configs{}, errors.New("get config failed")
	}
	var iniPath string
	dev := os.Getenv("Develop")
	if dev == "True" || dev == "1" {
		iniPath = cwdPath + "/conf/development.ini"
	}else {
		iniPath = cwdPath + "/conf/production.ini"
	}

	cfg, err := ini.Load(iniPath)
	if err != nil {
		// os.Exit(1)
		fmt.Println("获取配置文件失败...", err.Error())
		return model.Configs{}, errors.New("read config failed")
	}
	// 配置输出到结构体中
	conf = model.Configs{
		AppId: cfg.Section("server").Key("app_id").String(),
		Port: cfg.Section("server").Key("port").MustInt(8020),
		Mode: cfg.Section("server").Key("mode").String(),
		ReadTimeout: time.Duration(cfg.Section("server").Key("readtimeout").MustInt(60)) * time.Second,
		WriteTimeout: time.Duration(cfg.Section("server").Key("writetimeout").MustInt(60)) * time.Second,

		LogPath: cfg.Section("log").Key("logpath").String(),

		Referrer: cfg.Section("auth").Key("referrer").String(),
		CrsfToken: cfg.Section("auth").Key("crsftoken").String(),
		Host: cfg.Section("auth").Key("host").String(),
		TimeDiff: cfg.Section("auth").Key("timediff").MustInt64(60 * 5),
	}

	return conf, nil
}

func ReadAppInfo() (model.AppInfo, error) {
	var cwdPath string
	var info model.AppInfo
	cwdPath , pathErr := os.Getwd()
	if pathErr != nil {
		return model.AppInfo{}, errors.New("get appinfo failed")
	}
	iniPath := cwdPath + "/conf/app_info.ini"
	cfg, err := ini.Load(iniPath)
	if err != nil {
		// os.Exit(1)
		fmt.Println("获取配置文件失败...")
		return model.AppInfo{}, errors.New("read appinfo failed")
	}
	// 配置输出到结构体中
	info = model.AppInfo{
		App_id: cfg.Section("").Key("app_id").String(),
		App_des: cfg.Section("").Key("app_description").String(),
		App_version: cfg.Section("").Key("app_version").String(),
		App_build: cfg.Section("").Key("build_version").String(),
	}

	return info, nil
}