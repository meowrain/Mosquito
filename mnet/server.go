package mnet

import (
	"fmt"
	"mosquito/miface"
	. "mosquito/mlogger"
	"net"
)

// Server IServer的接口实现
type Server struct {
	Name      string //服务器的名称
	IPVersion string //服务器绑定的版本
	IP        string //服务器监听的ip
	Port      uint   //服务器监听的端口
}

func (server *Server) Start() {
	MLogger.Info(fmt.Sprintf("Starting server %s on %s:%d", server.Name, server.IP, server.Port))
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
		for {
			conn, err := listener.Accept()
			defer conn.Close()
			if err != nil {
				MLogger.Error(fmt.Sprintf("accept failed: %v", err))
				continue
			}
			MLogger.Info(fmt.Sprintf("Accepted connection from %s", conn.RemoteAddr()))
			go func() {
				for {
					buf := make([]byte, 1024)
					cnt, err := conn.Read(buf)
					if err != nil {
						MLogger.Error(fmt.Sprintf("receive from client failed: %v", err))
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						MLogger.Error(fmt.Sprintf("send to client failed: %v", err))
					}
				}
			}()
		}

	}()
}

func (server *Server) Stop() {
	//TODO implement me

}

func (server *Server) Serve() {
	server.Start()
	//做一些启动服务器后要做的其它工作
	//阻塞
	select {}
}

// 初始化Server模块
func NewServer(name string) miface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8099,
	}
	return s
}
