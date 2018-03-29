package config

import "github.com/bysir-zl/bygo/config"

var Debug = false
var DataSource = ""
var Listen = ""
var Key = ""
var WxAppId = ""

func init() {
	config.Load("app.ini")

	Debug = config.GetBool("debug", "")
	DataSource = config.GetString("datasource", "mysql")
	Listen = config.GetString("listen", "")
	Key = config.GetString("key", "")
	WxAppId = config.GetString("WxAPPID", "wx")
}

