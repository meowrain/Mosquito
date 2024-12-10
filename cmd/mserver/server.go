package main

import (
	"mosquito/conf"
	"mosquito/mnet"
)
import "mosquito/mlogger"

func main() {
	//初始化日志库
	mlogger.InitLogger()

	server := mnet.NewServer(conf.GlobalConf.App.Name)
	//添加路由
	server.AddRouter(&mnet.BaseRouter{})
	server.Serve()
}
