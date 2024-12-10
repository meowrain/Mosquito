package main

import (
	"fmt"
	"mosquito/conf"
	"mosquito/miface"
	"mosquito/mlogger"
	"mosquito/mnet"
)

type TestRouter struct {
	mnet.BaseRouter
}

func (br *TestRouter) Handle(request miface.IRequest) {
	msgId := request.GetMsgID()
	data := request.GetData()
	fmt.Printf("got msg from client: %v %v\n", msgId, string(data))
}
func main() {
	//初始化日志库
	mlogger.InitLogger()

	server := mnet.NewServer(conf.GlobalConf.App.Name)
	//添加路由
	server.AddRouter(&TestRouter{})
	server.Serve()
}
