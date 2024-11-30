package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("客户端启动....")
	fmt.Println("连接到: tcp://127.0.0.1:8099")
	conn, err := net.Dial("tcp4", "127.0.0.1:8099")
	if err != nil {
		fmt.Printf("连接失败:%v\n", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello world"))

	if err != nil {
		fmt.Printf("写入失败:%v\n", err)
		return
	}
	buf := make([]byte, 1024)
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("读取失败:%v\n", err)
		return
	}
	fmt.Println(string(buf[:cnt]))

}
