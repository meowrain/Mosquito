package mnet

import (
	"fmt"
	"io"
	"mosquito/miface"
	"mosquito/mlogger"
	"net"
)

type Connection struct {
	//当前链接的套接字
	Conn *net.TCPConn
	//链接的ID
	ConnID uint32

	//当前链接的状态
	isClosed bool

	//当前链接所绑定的业务方法API
	handleAPI miface.HandleFunc

	//告知当前链接已经退出/停止 的channel
	ExitChan chan bool

	//该链接处理的方法Router
	Router miface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router miface.IRouter) Connection {
	c := Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}
func (c *Connection) StartReader() {
	mlogger.MLogger.Info(fmt.Sprintf("Starting reader for connection id %d,received connection from %v\n", c.GetConnectionID(), c.GetTcpConnection()))
	defer c.Stop()
	for {
		buf := make([]byte, 1024)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				mlogger.MLogger.Info(fmt.Sprintf("Connection %v has been closed.", c.GetConnectionID()))
				return
			}
			mlogger.MLogger.Error(fmt.Sprintf("Error reading from connection id %d: %s\n", c.ConnID, err.Error()))
			continue
		}
		// 得到Request
		req := &Request{
			conn: c,
			data: buf[:cnt],
		}
		// 执行注册路由的方法
		go func(request miface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(req)

	}

}
func (c *Connection) StartWriter() {
	mlogger.MLogger.Info(fmt.Sprintf("Starting writer for connection id %d\n", c.ConnID))

}
func (c *Connection) Start() {
	mlogger.MLogger.Info(fmt.Sprintf("Starting connection %d\n", c.ConnID))
	//启动从当前链接读的业务
	go c.StartReader()
	//启动从当前链接写的业务
	go c.StartWriter()

}

func (c *Connection) Stop() {
	mlogger.MLogger.Info(fmt.Sprintf("Connection %d stop\n", c.ConnID))
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()

	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetConnectionID() uint32 {
	return c.ConnID
}

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	panic("implement me")
}
