package main

import (
	"gvb_server/core"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//	读取配置文件
	core.InitConf()
	//	初始化日志
	global.Log = core.InitLogger()
	//	连接数据库
	global.DB = core.InitGorm()
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
