/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-23
*/
package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"jjservice/src/util"
)

func initDefault(r *gin.Engine) {
	defaultHandler := func(c *gin.Context) {
		util.JJResponse(
			c,
			http.StatusOK,
			"default page",
			c.FullPath(),
		)
	}

	De := r.Group("")
	{
		De.GET("/", home)
		De.GET("/api", home)
		De.GET("/about", defaultHandler)
		De.GET("/api/logo", logo)
		De.POST("/api/login", login)
		De.POST("/api/changelog", changelog)
	}
}

// 主页，使用渲染模板的方式
func home(c *gin.Context) {
	util.JJResponse(
		c,
		http.StatusOK,
		"welcome to JJ Service",
		"Powered by app.renj.io",
		)
}

// 返回logo
func logo(c *gin.Context) {
	path, _ := os.Getwd()
	path = path + "/staticfile/logo.png"
	util.JJFile(
		c,
		path,
		)
}

func login(c *gin.Context) {
	// 验证来自参数的token
	inToken := c.Request.Header.Get("tokenstring")
	newToken := ""
	if inToken != "" && inToken == util.Conf.CrsfToken {
		newToken = util.TokenEncrypt()
	}

	util.JJResponse(
		c,
		http.StatusOK,
		"Login to JJ Service",
		newToken,
	)
}

func changelog(c *gin.Context) {
	cwd, _ := os.Getwd()
	var logE  []interface{}
	data, err := ioutil.ReadFile(path.Join(cwd, "conf", "update.json"))
	if err != nil {
		util.JJResponse(
			c,
			http.StatusOK,
			"This is jjservice's change logs.",
			"",
		)
	}else {
		json.Unmarshal(data, &logE)
		util.JJResponse(
			c,
			http.StatusOK,
			"This is jjservice's change logs.",
			logE,
		)
	}

}