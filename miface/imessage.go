package miface

type IMessage interface {
	GetMsgID() uint32    //获取消息ID
	GetMsgLen() uint32   //获取消息长度
	GetData() []byte     //获取消息内容
	SetMsgID(id uint32)  //设置消息ID
	SetMsgLen(l uint32)  //设置消息长度
	SetData(data []byte) //设置消息内容
}
