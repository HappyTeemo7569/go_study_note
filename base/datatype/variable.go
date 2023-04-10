package main

import (
	"fmt"
	"reflect"
)

//注意：自动推导类型只能在函数内部使用
//v4 := 2 //err

// 变量声明
func main() {
	var v1 int = 10 // 方式1
	var v2 = 10     // 方式2，编译器自动推导出v2的类型
	v3 := 10        // 方式3，编译器自动推导出v3的类型

	fmt.Println("v1 type is ", reflect.TypeOf(v1))
	fmt.Println("v2 type is ", reflect.TypeOf(v2))
	fmt.Println("v3 type is ", reflect.TypeOf(v3)) //v3 type is  int

	//自动推导类型
	//出现在 := 左侧的变量不应该是已经被声明过，:=定义时必须初始化
	var v4 int
	//v4 := 2 //err 重复定义

	fmt.Println("v4 type is ", reflect.TypeOf(v4))
}

// 变量赋值
func setValue() {
	var v1 int
	v1 = 123

	var v2, v3, v4 int
	v2, v3, v4 = 1, 2, 3 //多重赋值

	fmt.Println(v1, v2, v3, v4)

	i := 10
	j := 20
	i, j = j, i //多重赋值
}

func test() (int, string) {
	return 250, "sb"
}

// 匿名变量
func anonymous() {
	_, i, _, j := 1, 2, 3, 4
	fmt.Println(i, j)

	_, str := test()
	fmt.Println(str)
}
