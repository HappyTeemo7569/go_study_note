package main

import "fmt"

func test(x int) {
	defer fmt.Println("我要崩溃了")
	fmt.Println(100 / x) //x为0时，产生异常
}

func main() {
	defer func() {
		recover()
		fmt.Println("没事，我兜住了")
	}()

	defer test(0)

	defer fmt.Println("aaaaaaaa")
	defer fmt.Println("bbbbbbbb")

	/*
		   运行结果：
		    bbbbbbbb
			aaaaaaaa
			我要崩溃了
			没事，我兜住了
	*/
}
