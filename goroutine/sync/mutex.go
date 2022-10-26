package main

import (
	"fmt"
	"sync"
	"time"
)

/**
共享内存
*/

var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.Mutex{} //初始化锁

	for i := 0; i < 100; i++ {
		go add(1, i, lock) //并发相加  注意用同一个锁
	}

	//for {
	//	lock.Lock()
	//	c := counter
	//	lock.Unlock()
	//	println("循环")
	//	time.Sleep(1 * time.Second)
	//	//runtime.Gosched() //让出CPU时间片
	//	if c >= 10 {
	//		break
	//	}
	//}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
