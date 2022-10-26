package main

/**
其主要用途是保证指定函数代码只执行一次，类似于单例模式，常用于应用启动时的一些全局初始化操作。
它只提供了一个 Do 方法，该方法只接受一个参数，且这个参数的类型必须是 func()，即无参数无返回值的函数类型。
*/

import (
	"fmt"
	"sync"
	"time"
)

func dosomething(o *sync.Once) {
	fmt.Println("Start:")
	o.Do(func() {
		fmt.Println("Do Something...")
	})
	fmt.Println("Finished.")
}

func main() {
	o := &sync.Once{}
	go dosomething(o)
	go dosomething(o)
	time.Sleep(time.Second * 1)
}
