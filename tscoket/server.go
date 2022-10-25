package tscoket

import (
	"fmt"
	"net"
	"os"
)

func TCPServerRun() {

	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭监听
	defer listener.Close()

	fmt.Println("TCP开始监听")

	//阻塞等待用户链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//接收用户的请求
	buf := make([]byte, 1024) //1024大小的缓冲区
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Println("收到消息:buf = ", string(buf[:n]))

	defer conn.Close() //关闭当前用户链接
}

func UDPServerRun() {

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}

	//监听
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("UPP开始监听")

	//关闭监听
	defer listener.Close()

	//接收用户的请求
	for {
		data := make([]byte, 1024)
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println("failed to read UDP msg because of ", err.Error())
			return
		}
		fmt.Println(n, remoteAddr)              //长度和IP
		fmt.Println("buf = ", string(data[:n])) //打印内容
	}
}
