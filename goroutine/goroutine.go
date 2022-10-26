package main

import (
	"fmt"
	"runtime"
	"time"
)

func task(g_index int) {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine:index:%d : i = %d \n", g_index, i)
		time.Sleep(1 * time.Second) //延时1s
	}
}

func main() {

	for g_index := 0; g_index <= 10; g_index++ {
		go task(g_index)
		fmt.Printf(" goroutine num = %d\n", runtime.NumGoroutine())
	}

	i := 0
	//main goroutine 循环打印
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}

}
