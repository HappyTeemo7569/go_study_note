package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add_c(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
}

func main() {
	/**
	select {
	    case <-chan1:
	        // 如果从 chan1 通道成功接收数据，则执行该分支代码
	    case chan2 <- 1:
	        // 如果成功向 chan2 通道成功发送数据，则执行该分支代码
	    default:
	        // 如果上面都没有成功，则进入 default 分支处理流程
	}
	*/
	chs := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3) // 随机生成0-2之间的数字
	fmt.Printf("随机索引/数值: %d\n", index)
	chs[index] <- index // 向通道发送随机数字

	// 哪一个通道中有值，哪个对应的分支就会被执行
	select {
	case <-chs[0]:
		fmt.Println("第一个条件分支被选中")
	case <-chs[1]:
		fmt.Println("第二个条件分支被选中")
	case num := <-chs[2]:
		fmt.Println("第三个条件分支被选中:", num)
	default:
		fmt.Println("没有分支被选中")
	}
}

func RunC() {
	start := time.Now()
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go add_c(1, i, chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)

}
