package mnet

import (
	"errors"
	"fmt"
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
		//创建数据包对象
		dp := NewDataPack()
		//读取包头 包括id和数据长度
		headData := make([]byte, dp.GetHeadLen())
		if _, err := c.Conn.Read(headData); err != nil {
			break
		}
		//把包头信息解包到message结构体中
		message, err := dp.Unpack(headData)
		if err != nil {
			mlogger.MLogger.Error(err.Error())
			break
		}
		//从message中拿到包头大小
		bodyLen := message.GetMsgLen()
		//读取客户端发送的data到bodyData切片中,如果数据长度>0就可以存了
		if bodyLen > 0 {
			bodyData := make([]byte, bodyLen)
			if _, err := c.Conn.Read(bodyData); err != nil {
				mlogger.MLogger.Error(err.Error())
				break
			}
			//存储到message结构体
			message.SetData(bodyData)
		}
		//构建Request结构体对象
		req := &Request{
			conn: c,
			msg:  message,
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

func (c *Connection) Send(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection is closed")
	}
	msg := NewMessagePackage(msgId, data)
	dp := NewDataPack()
	bytesData, err := dp.Pack(msg)
	if err != nil {
		mlogger.MLogger.Error(err.Error())
		return errors.New("Pack Error msg")
	}
	if _, err := c.Conn.Write(bytesData); err != nil {
		mlogger.MLogger.Error(fmt.Sprintf("Write Error msg,msg id : %v,err: %v", msg.GetMsgID(), err))
		return errors.New(fmt.Sprintf("Write Error msg,msg id : %v", msg.GetMsgID()))
	}
	return nil
}
