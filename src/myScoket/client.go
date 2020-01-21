package myScoket

import (
	"fmt"
	"net"
	"os"
)

func TCPClient() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//延时关闭
	defer conn.Close()

	//发送数据
	conn.Write([]byte("你好"))
	fmt.Println("发送消息完毕")

}

func UDPClient() {
	//主动连接服务器
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}

	//延时关闭
	defer conn.Close()

	//发送数据
	conn.Write([]byte("你好"))
	fmt.Println("UDP发送消息完毕")

}
