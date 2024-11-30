package miface

type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
	// AddRouter 路由： 给当前的服务注册一个路由方法，供对客户端的Request进行处理
	AddRouter(router IRouter)
}
