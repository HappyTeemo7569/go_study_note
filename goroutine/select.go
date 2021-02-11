package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
select {
    case <-chan1:
        // 如果从 chan1 通道成功接收数据，则执行该分支代码
    case chan2 <- 1:
        // 如果成功向 chan2 通道成功发送数据，则执行该分支代码
    default:
        // 如果上面都没有成功，则进入 default 分支处理流程
}

select 不像 switch，case 后面并不带判断条件，而是直接去查看 case 语句，每个 case 语句都必须是一个面向通道的操作。
需要注意的是这两个 case 的执行不是 if...else... 那种先后关系，而是会并发执行.
如果两者同时返回，则随机选择一个执行，如果这两者都没有返回，则进入 default 分支，这里也不会出现阻塞，如果 chan1 通道为空，或者 chan2 通道已满，就会立即进入 default 分支
如果没有 default 语句，则会阻塞直到某个通道操作成功。
*/

func main1() {
	rand.Seed(time.Now().UnixNano())

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

//select 语句只能对其中的每一个 case 表达式各求值一次，如果我们想连续操作其中的通道的话，需要通过在 for 语句中嵌入 select 语句的方式来实现

func main() {
	rand.Seed(time.Now().UnixNano())

	chs := [3]chan int{
		make(chan int, 3),
		make(chan int, 3),
		make(chan int, 3),
	}

	index1 := rand.Intn(3) // 随机生成0-2之间的数字
	fmt.Printf("随机索引/数值: %d\n", index1)
	chs[index1] <- rand.Int() // 向通道发送随机数字

	index2 := rand.Intn(3)
	fmt.Printf("随机索引/数值: %d\n", index2)
	chs[index2] <- rand.Int()

	index3 := rand.Intn(3)
	fmt.Printf("随机索引/数值: %d\n", index3)
	chs[index3] <- rand.Int()

	// 哪一个通道中有值，哪个对应的分支就会被执行
	for i := 0; i < 3; i++ {
		select {
		case num, ok := <-chs[0]:
			if !ok {
				break
			}
			fmt.Println("第一个条件分支被选中: chs[0]=>", num)
		case num, ok := <-chs[1]:
			if !ok {
				break
			}
			fmt.Println("第二个条件分支被选中: chs[1]=>", num)
		case num, ok := <-chs[2]:
			if !ok {
				break
			}
			fmt.Println("第三个条件分支被选中: chs[2]=>", num)
		default:
			fmt.Println("没有分支被选中")
		}
	}
}
