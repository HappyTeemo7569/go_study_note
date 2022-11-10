package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

/*
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
Deadline：
返回值1：截止时间，到了这个时间点，Context 会自动触发 Cancel 动作。
返回值2：一个布尔值，true 表示设置了截止时间，false 表示没有设置截止时间，如果没有设置截止时间，就要手动调用 cancel 函数取消 Context。

Done：
返回一个只读的通道（只有在被cancel后才会返回），类型为 struct{}。
当这个通道可读时，意味着parent context已经发起了取消请求，根据这个信号，开发者就可以做一些清理动作，退出goroutine。

Err：
回 context 被 cancel 的原因。

Value：
返回被绑定到 Context 的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。

*/

/**
当一个协程（goroutine）开启后，我们是无法强制关闭它的。
常见的关闭协程的原因有如下几种：
1.goroutine 自己跑完结束退出
2.主进程crash退出，goroutine 被迫退出
3.通过通道发送信号，引导协程的关闭。

第一种，属于正常关闭，不在今天讨论范围之内。
第二种，属于异常关闭，应当优化代码。
第三种，才是开发者可以手动控制协程的方法
*/

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

/**

使用content注意：
1. 通常 Context 都是做为函数的第一个参数进行传递（规范性做法），并且变量名建议统一叫 ctx
2.Context 是线程安全的，可以放心地在多个 goroutine 中使用。
3.当你把 Context 传递给多个 goroutine 使用时，只要执行一次 cancel 操作，所有的 goroutine 就可以收到 取消的信号
4.不要把原本可以由函数参数来传递的变量，交给 Context 的 Value 来传递。
5.当一个函数需要接收一个 Context 时，但是此时你还不知道要传递什么 Context 时，可以先用 context.TODO 来代替，而不要选择传递一个 nil。
6.当一个 Context 被 cancel 时，继承自该 Context 的所有 子 Context 都会被 cancel。

*/
