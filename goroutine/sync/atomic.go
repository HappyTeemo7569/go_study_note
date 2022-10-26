package main

/**
竞态检测: 检测代码在并发环境下可能出现的问题
当多线程并发运行的程序竞争访问和修改同块资源时，会发生竞态问题。

这些原子操作包括加法（Add）、比较并交换（Compare And Swap，简称 CAS）、加载（Load）、存储（Store）和交换（Swap）
*/

import (
	"fmt"
	"sync/atomic"
)

var (
	// 序列号
	seq int64
)

// 序列号生成器
func GenID() int64 {

	// 尝试原子的增加序列号
	atomic.AddInt64(&seq, 1)
	return seq
}

/**
go run -race  可以竞态检测。
*/
func main() {

	// 10个并发序列号生成
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())
}
