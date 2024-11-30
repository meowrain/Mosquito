package mnet

import "mosquito/miface"

type Request struct {
	//客户端与服务器的连接
	conn miface.IConnection
	// 客户端请求的数据
	data []byte
}
	
func (r *Request) GetConnection() miface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
