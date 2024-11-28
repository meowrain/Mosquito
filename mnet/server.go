package mnet

import "mosquito/miface"

// Server IServer的接口实现
type Server struct {
	Name      string //服务器的名称
	IPVersion string //服务器绑定的版本
	IP        string //服务器监听的ip
	Port      uint   //服务器监听的端口
}

func (server *Server) Start() {
	//TODO implement me

}

func (server *Server) Stop() {
	//TODO implement me

}

func (server *Server) Serve() {
	//TODO implement me

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
