package main

import (
	"fmt"
	"time"
)

func add2(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1 //执行完了就写一条表示自己完成了
}

func main() {
	start := time.Now()
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add2(1, i, chs[i]) //分配了10个协程出去了
	}
	for _, ch := range chs {
		<-ch //循环等待，要每个完成才能继续，不然就等待
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

/**
在每个协程的 add() 函数业务逻辑完成后，我们通过 ch <- 1 语句向对应的通道中发送一个数据。
在所有的协程启动完成后，我们再通过 <-ch 语句从通道切片 chs 中依次接收数据（不对结果做任何处理，相当于写入通道的数据只是个标识而已，表示这个通道所属的协程逻辑执行完毕），
直到所有通道数据接收完毕，然后打印主程序耗时并退出。
*/
