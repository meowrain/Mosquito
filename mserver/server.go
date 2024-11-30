package main

import "mosquito/mnet"
import "mosquito/mlogger"

func main() {
	//初始化日志库
	mlogger.InitLogger()

	server := mnet.NewServer("Mosquito")
	server.Serve()
}
