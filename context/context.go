package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func AddNum(a *int32, b int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(a)
		newNum := curNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(a, curNum, newNum) {
			fmt.Printf("number当前值: %d [%d-%d]\n", *a, b, i)
			break
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", b, i)
		}
	}
}

func main() {
	total := 10
	var num int32
	fmt.Printf("number初始值: %d\n", num)
	fmt.Println("启动子协程...")
	ctx, cancelFunc := context.WithCancel(context.Background())

	/**
	先通过 context.WithCancel 方法返回一个新的 cxt 和 cancelFunc，
	并且通过 context.Background() 方法传入父 Context，
	该 Context 没有值，永远不会取消，可以看作是所有 Context 的根节点，比如这里的 cxt 就是从父 Context 拷贝过来的可撤销的子 Context。
	*/

	for i := 0; i < total; i++ {
		go AddNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) { //（所有子协程执行完毕）
				cancelFunc() //撤销对应子 Context 对象 cxt，这样，处于阻塞状态的 cxt.Done() 对应通道被关闭
			}
		})
	}
	<-ctx.Done()

	/**
	注：cxt.Done() 方法返回一个通道，该通道会在调用 cancelFunc 函数时关闭，或者在父 context 撤销时也会被关闭。
	*/

	fmt.Println("所有子协程执行完毕.")
}

/**
WithDeadline 和 WithTimeout 分别比 WithCancel 多了一个 deadline 和 timeout 时间参数，
表示子 Context 存活的最长时间，如果超过了该时间，会自动撤销对应的子 Context。
相应的，在调用 <-cxt.Done() 等待子协程执行结束时，如果没有调用 cancelFunc 函数的话它们会等待过期时间到达自动关闭，
不过我们通常还是会主动调用 cancelFunc 函数以便更好的控制程序运行。
*/

func test() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	valueCtx := context.WithValue(ctx, "key", "value")
	defer cancelFunc()
	for i := 0; i < total; i++ {
		go AddNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				fmt.Println("key:", valueCtx.Value("key"))
				cancelFunc()
			}
		})
	}
	<-ctx.Done()
}
