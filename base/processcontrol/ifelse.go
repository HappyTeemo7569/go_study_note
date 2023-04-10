package main

import (
	"fmt"
	"time"
)

func main() {

	a := time.Now().Unix() / 2

	if a == 2 { //条件表达式没有括号
		fmt.Println("a==2")
	}

	//支持一个初始化表达式, 初始化字句和条件表达式直接需要用分号分隔
	if b := time.Now().Unix() / 3; b == 3 {
		fmt.Println("b==3")
	}
	/**
	Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch
	但是可以使用fallthrough强制执行后面的case代码：
	*/

	var score int = 90

	switch score {
	case 90:
		fmt.Println("优秀")
	//fallthrough
	case 80:
		fmt.Println("良好")
	//fallthrough
	case 50, 60, 70: //一个case多个条件
		fmt.Println("一般")
	//fallthrough
	default:
		fmt.Println("差")
	}
}
