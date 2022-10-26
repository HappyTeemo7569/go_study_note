package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	index := 1
	fmt.Println("协程创建了")
	for {
		time.Sleep(time.Second * 2) //延时2s
		fmt.Println("协程 index", index)
		index++
	}
}

func main() {

	go newTask() //新建一个协程， 新建一个任务
	index := 1
	//死循环
	for {
		time.Sleep(time.Second) //延时1s

		runtime.Gosched() //让别人先执行，需要同时要时间片的时候才会有效，对方如果已经停了就还是自己执行。

		fmt.Println("主进程 index", index)
		index++
	}

	// 主goroutine退出后，其它的工作goroutine也会自动退出
	// 多个 goroutine 的执行顺序不定
}
