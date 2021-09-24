package main

import (
	"fmt"
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

func Run() {

	go newTask() //新建一个协程， 新建一个任务
	index := 1
	//死循环
	for {
		time.Sleep(time.Second) //延时1s
		fmt.Println("主进程 index", index)
		index++
	}
}
