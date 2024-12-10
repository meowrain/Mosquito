package mnet

import (
	"fmt"
	"net"
	"testing"
)

func TestDataPackServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept err:", err)
				return
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := conn.Read(headData)
					if err != nil {
						break
					}
					headMsg, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("unpack head err:", err)
						return
					}
					dataLen := headMsg.GetMsgLen()
					msgId := headMsg.GetMsgID()
					if dataLen < 0 {
						fmt.Println("无数据")
						return
					}
					bodyData := make([]byte, dataLen)
					_, err = conn.Read(bodyData)
					if err != nil {
						fmt.Println("read body err:", err)
						return
					}
					msg := &Message{
						Id:      msgId,
						DataLen: dataLen,
						Data:    bodyData,
					}
					fmt.Println("接收到的msg信息：", msg)
				}
			}(conn)
		}
	}()
	select {}
}

func TestDatapackClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":7777")
	if err != nil {
		fmt.Println("server connect err:", err)
	}
	dp := NewDataPack()
	data := "hello world"
	msg1 := Message{
		Id:      1,
		DataLen: uint32(len(data)),
		Data:    []byte(data),
	}
	msg2 := Message{
		Id:      2,
		DataLen: uint32(len(data)),
		Data:    []byte(data),
	}

	pack1, err := dp.Pack(&msg1)
	if err != nil {
		fmt.Println("pack err:", err)
		return
	}
	pack2, err := dp.Pack(&msg2)
	if err != nil {
		fmt.Println("pack err:", err)
		return
	}
	pack1 = append(pack1, pack2...)
	_, err = conn.Write(pack1)
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}
}
