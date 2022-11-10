package main

import "fmt"

/**
信道死锁
*/

//对于无缓冲信道，在接收者未准备好之前，发送操作是阻塞的.
func question1() {
	pipline := make(chan string)
	pipline <- "hello world" //fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-pipline)
	//对于解决此问题有两种方法：
	//	1.
	//使接收者代码在发送者之前执行
	//	2.
	//使用缓冲信道，而不使用无缓冲信道
}
func hello(pipline chan string) {
	<-pipline
}

func solution1() {
	pipline := make(chan string)
	go hello(pipline) //不异步仍然会阻塞
	pipline <- "hello world"
}

func solution2() {
	pipline := make(chan string, 1)
	pipline <- "hello world"
	fmt.Println(<-pipline)
}

func main() {
	//question1()
	solution1()
	solution2()

	//question2()
	question3()
}

//信道容量为 1，但是往信道中写入两条数据，对于一个协程来说就会造成死锁。
func question2() {
	ch1 := make(chan string, 1)

	ch1 <- "hello world"
	ch1 <- "hello China"

	fmt.Println(<-ch1)
}

//当程序一直在等待从信道里读取数据，而此时并没有人会往信道中写入数据。此时程序就会陷入死循环，造成死锁。
func question3() {
	pipline := make(chan string)
	go func() {
		pipline <- "hello world"
		pipline <- "hello China"
		// close(pipline)  //放开这句即可
	}()
	for data := range pipline {
		fmt.Println(data)
	}
}
