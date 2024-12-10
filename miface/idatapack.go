package miface

type IDataPack interface {
	//GetHeadLen 获取包的长度的方法
	GetHeadLen() uint32
	//Pack 封包方法
	Pack(msg IMessage) ([]byte, error)
	//Unpack 拆包方法
	Unpack([]byte) (IMessage, error)
}
