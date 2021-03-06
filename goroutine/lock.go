package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter_r int = 0

func add_r(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter_r++
	fmt.Printf("%d: %d + %d = %d\n", counter_r, a, b, c)
	lock.Unlock()
}

func rLockTest() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add_r(1, i, lock)
	}

	for {
		lock.RLock()
		c := counter_r
		lock.RUnlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}

func RunL() {
	rLockTest()
}
