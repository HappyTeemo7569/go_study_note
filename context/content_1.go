package main

import (
	"context"
	"fmt"
	"time"
)

/**
chan+select的方式，是比较优雅的结束一个goroutine的方式.
不过这种方式也有局限性，如果有很多goroutine都需要控制结束怎么办呢？
如果这些goroutine又衍生了其他更多的goroutine怎么办呢？如果一层层的无穷尽的goroutine呢？
这就非常复杂了，即使我们定义很多chan也很难解决这个问题，因为goroutine的关系链就导致了这种场景非常复杂。
*/

func stopBySelect() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}

func monitor(ctx context.Context, number int) {
	for {
		select {
		// 其实可以写成 case <- ctx.Done()
		// 这里仅是为了让你看到 Done 返回的内容
		case v := <-ctx.Done():
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束。\n", number, v)
			return
		default:
			fmt.Printf("监控器%v，正在监控中...\n", number)
			time.Sleep(2 * time.Second)
		}
	}
}

func stopByContent() {
	ctx, cancel := context.WithCancel(context.Background())

	//添加截止时间 传入绝对时间
	//ctx02, cancel := context.WithDeadline(ctx01, time.Now().Add(1 * time.Second))
	//添加超时 传入相对时间
	//ctx02, cancel := context.WithTimeout(ctx01, 1* time.Second)
	//添加参数
	//ctx03 := context.WithValue(ctx02, "item", "CPU")
	//value := ctx.Value("item")

	for i := 1; i <= 5; i++ {
		go monitor(ctx, i)
	}

	time.Sleep(1 * time.Second)
	// 关闭所有 goroutine
	cancel()

	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)

	fmt.Println("主程序退出！！")
}

func main() {
	stopByContent()
}
