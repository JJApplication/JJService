/*
App: JJService
Author: Landers
Github: https://github.com/landers1037
Date: 2020-09-24
*/
package model

// 常量包 转为使用json的方式更新
type Apps struct {
	App_id 		string
	App_name 	string
	App_code    string
	App_des 	string
	App_version string
	App_build   string
	App_port 	int
	App_dev 	bool // 是否开发完成
}


var ALL_APPS = map[string]Apps{}

// 舍弃在go中定义 改为通过json文件加载
func init() {
	ALL_APPS = make(map[string]Apps)
	ALL_APPS["app_jjservice"] = Apps{
			App_id:      "app_jjservice",
			App_name:    "JJService",
			App_code:    "jjservice",
			App_des:     "服务管理平台 For Renj.io",
			App_version: "1.3",
			App_build:   "build2020.10.1",
			App_port:    8004,
			App_dev:     true,
		}
	ALL_APPS["app_jjgo"] = Apps{
		App_id:      "app_jjgo",
		App_name:    "JJGo",
		App_code:    "jjgo",
		App_des:     "JJ Restful API接口",
		App_version: "1.0",
		App_build:   "build2020.9.30",
		App_port:    8007,
		App_dev:     false,
	}
	ALL_APPS["app_jjbox"] = Apps{
		App_id:      "app_jjbox",
		App_name:    "JJBox",
		App_code:    "jjbox",
		App_des:     "在线云盘",
		App_version: "1.0",
		App_build:   "build2020.9.30",
		App_port:    0,
		App_dev:     false,
	}
	ALL_APPS["app_mysite"] = Apps{
		App_id:      "app_mysite",
		App_name:    "Renj.io",
		App_code:    "mysite",
		App_des:     "我的网站主页",
		App_version: "4.2",
		App_build:   "build2020.9.26",
		App_port:    8000,
		App_dev:     true,
	}
	ALL_APPS["app_blog"] = Apps{
		App_id:      "app_blog",
		App_name:    "Vue Blog",
		App_code:    "blog",
		App_des:     "前后端分离设计的Vue轻量博客",
		App_version: "1.23",
		App_build:   "build2020.8.5",
		App_port:    8003,
		App_dev:     true,
	}
	ALL_APPS["app_celery"] = Apps{
		App_id:      "app_celery",
		App_name:    "Auto Mail Service",
		App_code:    "celery",
		App_des:     "自动化邮件订阅服务",
		App_version: "1.4",
		App_build:   "build2020.8.30",
		App_port:    0,
		App_dev:     true,
	}
	ALL_APPS["app_mgek"] = Apps{
		App_id:      "app_mgek",
		App_name:    "Mgek APP",
		App_code:    "mgek",
		App_des:     "Mgek项目",
		App_version: "4.11",
		App_build:   "build2020.9.9",
		App_port:    8010,
		App_dev:     true,
	}
	ALL_APPS["app_mgekfile"] = Apps{
		App_id:      "app_mgekfile",
		App_name:    "Mgek File",
		App_code:    "mgekfile",
		App_des:     "Mgek缓存服务",
		App_version: "1.3",
		App_build:   "build2020.4.20",
		App_port:    8011,
		App_dev:     true,
	}
	ALL_APPS["app_mgekdoc"] = Apps{
		App_id:      "app_mgekdoc",
		App_name:    "Mgek Doc",
		App_code:    "mgekdoc",
		App_des:     "Mgek在线文档",
		App_version: "2.4",
		App_build:   "build2020.8.20",
		App_port:    8010,
		App_dev:     true,
	}
}



