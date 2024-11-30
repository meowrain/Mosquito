package miface

/*
IRequest 接口把客户端请求的链接和请求的数据绑定在一起
*/
type IRequest interface {
	//获取当前连接
	GetConnection() IConnection
	//得到请求数据
	GetData() []byte
}
