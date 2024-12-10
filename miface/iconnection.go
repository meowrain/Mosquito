package miface

import "net"

type IConnection interface {
	// Start 启动该链接
	Start()
	// Stop 关闭链接
	Stop()
	// GetTcpConnection 获取链接
	GetTcpConnection() *net.TCPConn
	// GetRemoteAddr 获取客户端地址
	GetRemoteAddr() net.Addr
	// GetConnectionID 获取链接的ID
	GetConnectionID() uint32
	// Send 发送信息
	Send(uint32, []byte) error
}

// 定义一个处理链接业务的方法
// 参数分别为： 处理的链接，处理的数据，处理数据的长度
type HandleFunc func(*net.TCPConn, []byte, int) error
