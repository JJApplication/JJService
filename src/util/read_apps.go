/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-10-01
*/
package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"jjservice/src/model"
)

// 通过app.json文件导入apps信息
var ALL_APPS map[string]model.Apps

func init()  {
	_,_ = ReadApps()
}

func ReadApps() (map[string]model.Apps, error) {
	apps := make(map[string]model.Apps)
	p, _ := os.Getwd()
	localPath := path.Join(p, "conf", "apps.json")
	localAppsFile, err := ioutil.ReadFile(localPath)
	if err != nil {
		return apps, err
	}else {
		_ = json.Unmarshal(localAppsFile, &ALL_APPS)
		return ALL_APPS, nil
	}
}
