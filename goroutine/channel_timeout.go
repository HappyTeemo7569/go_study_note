package main

import (
	"fmt"
	"time"
)

func main() {
	// 初始化 ch 通道
	ch := make(chan int, 1)

	// 初始化 timeout 通道
	timeout := make(chan bool, 1)

	// 实现一个匿名超时等待函数
	go func() {
		time.Sleep(1e9) // 睡眠1秒钟
		timeout <- true
	}()

	// 借助 timeout 通道结合 select 语句实现 ch 通道读取超时效果
	select {
	case <-ch:
		fmt.Println("接收到 ch 通道数据")
	case <-timeout:
		fmt.Println("超时1秒，程序退出")
	}
}

/*
而如果没有 timeout 通道和上述 select 机制，从 ch 通道接收数据会得到如下 panic（死锁）：

fatal error: all goroutines are asleep - deadlock!

*/

//添加判断
func main2() {
	ch := make(chan int, 2)
	// 发送方
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("发送方: 发送数据 %v...\n", i)
			ch <- i
		}
		fmt.Println("发送方: 关闭通道...")
		close(ch)
	}()
	// 接收方
	for {
		num, ok := <-ch
		if !ok {
			fmt.Println("接收方: 通道已关闭")
			break
		}
		fmt.Printf("接收方: 接收数据: %v\n", num)
	}
	fmt.Println("程序退出")
}

/**
如果我们试图在通道 ch 关闭后发送数据到该通道，则会得到如下 panic：

panic: send on closed channel

而如果我们试图在通道 ch 关闭后再次关闭它，则会得到如下 panic：

panic: close of closed channel
*/
