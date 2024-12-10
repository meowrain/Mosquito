package mnet

import (
	"fmt"
	"mosquito/miface"
)

type Message struct {
	Id      uint32 //消息ID
	DataLen uint32 //消息长度
	Data    []byte //消息的内容
}

func NewMessagePackage(id uint32, data []byte) miface.IMessage {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}
func (m *Message) GetMsgID() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(id uint32) {
	m.Id = id
}

func (m *Message) SetMsgLen(datalen uint32) {
	m.DataLen = datalen
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) String() string {
	return fmt.Sprintf("id:%d data:%s datalen:%d", m.Id, string(m.Data), m.DataLen)
}
