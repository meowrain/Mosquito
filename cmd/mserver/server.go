package main

import "mosquito/mnet"
import "mosquito/mlogger"

func main() {
	//初始化日志库
	mlogger.InitLogger()

	server := mnet.NewServer("Mosquito")
	//添加路由
	server.AddRouter(&mnet.BaseRouter{})
	server.Serve()
}
