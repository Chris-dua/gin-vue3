package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title gvb_server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//	读取配置文件
	core.InitConf()
	//	初始化日志
	global.Log = core.InitLogger()
	//	连接数据库
	global.DB = core.InitGorm()
	// 连接redis
	global.Redis = core.ConnectRedis()
	global.ESClient = core.EsConnect()
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	addr := global.Config.System.Addr()
	router := routers.InitRouter()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
