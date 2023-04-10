package main

import "fmt"

/*
Go语言中不允许隐式转换，所有类型转换必须显式声明，而且转换只能发生在两种相互兼容的类型之间。
*/

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	//bool类型不能转换为int
	//fmt.Printf("flag = %d\n", int(flag))

	//0就是假，非0就是真
	//整型也不能转换为bool
	//flag = bool(1)

	var ch byte
	ch = 'a' //字符类型本质上就是整型
	var t int
	t = int(ch) //类型转换，把ch的值取出来后，转成int再给t赋值
	fmt.Println("t = ", t)
	//t =  97
}
