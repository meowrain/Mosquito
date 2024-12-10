package mnet

import (
	"fmt"
	"mosquito/conf"
	"mosquito/miface"
	. "mosquito/mlogger"
	"net"
	"os"
)

// Server IServer的接口实现
type Server struct {
	Name      string         //服务器的名称
	Version   string         //服务器版本
	IPVersion string         //服务器绑定的版本
	IP        string         //服务器监听的ip
	Port      uint           //服务器监听的端口
	Router    miface.IRouter //当前server的router
}

func (server *Server) Start() {
	MLogger.Info(fmt.Sprintf("Starting server %s-%s on %s:%d", server.Name, server.Version, server.IP, server.Port))
	go func() {

		//获取TCP ADDR
		addr, err := net.ResolveTCPAddr(server.IPVersion, fmt.Sprintf("%s:%d", server.IP, server.Port))
		if err != nil {
			MLogger.Error(fmt.Sprintf("resolve tcp addr failed: %v", err))
			return
		}
		//监听服务器地址
		listener, err := net.ListenTCP(server.IPVersion, addr)
		if err != nil {
			MLogger.Error(fmt.Sprintf("listen tcp failed: %v", err))
			return
		}
		MLogger.Info(fmt.Sprintf("%s Listening on %s:%d", server.Name, server.IP, server.Port))
		//阻塞等待客户端连接,处理相关请求
		var cid uint32 = 0
		for {
			conn, err := listener.AcceptTCP()
			defer conn.Close()
			if err != nil {
				MLogger.Error(fmt.Sprintf("accept failed: %v", err))
				continue
			}
			MLogger.Info(fmt.Sprintf("Accepted connection from %s", conn.RemoteAddr()))
			connectionObject := NewConnection(conn, cid, server.Router)

			connectionObject.Start()
			cid++
		}

	}()
}

func (server *Server) Stop() {
	os.Exit(1)
}

func (server *Server) Serve() {
	server.Start()
	//做一些启动服务器后要做的其它工作
	//阻塞
	select {}
}
func (server *Server) AddRouter(router miface.IRouter) {
	server.Router = router
}

// NewServer 初始化Server模块
func NewServer(name string) miface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		Version:   conf.GlobalConf.App.Version,
		IP:        conf.GlobalConf.App.Host,
		Port:      conf.GlobalConf.App.Port,
		Router:    nil,
	}
	return s
}
