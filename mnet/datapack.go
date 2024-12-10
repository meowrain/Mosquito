package mnet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"mosquito/conf"
	"mosquito/miface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 固定8字节
func (d *DataPack) GetHeadLen() uint32 {
	//Datalen uint32 --> 4字节  Id uint32 --> 4字节 加起来8字节
	return 8
}

func (d *DataPack) Pack(msg miface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	//将dataLen写入buff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//将MsgID写入buff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	//将Data数据写入buff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (d *DataPack) Unpack(binaryData []byte) (miface.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	//读datalen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//读msgid
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if conf.GlobalConf.App.MaxPackageSize > 0 && msg.DataLen > conf.GlobalConf.App.MaxPackageSize {
		return nil, fmt.Errorf("too large msg data recv")
	}
	return msg, nil
}
