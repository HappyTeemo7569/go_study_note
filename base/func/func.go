package main

import "fmt"

// 无参无返回值函数定义
func Test() {
	fmt.Println("this is a test func")
}

// 有参无返回值
func Test01(v1 int, v2 int) { //方式1
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)
}

func Test02(v1, v2 int) { //方式2, v1, v2都是int类型
	fmt.Printf("v1 = %d, v2 = %d\n", v1, v2)
}

// 不定参数
// 形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数
func Test03(args ...int) {
	for _, n := range args { //遍历参数列表
		fmt.Println(n)
	}
}

// 不定参数的传递
func Test04(args ...int) {
	Test03(args...)     //按原样传递, Test()的参数原封不动传递给Test03
	Test03(args[1:]...) //Test()参数列表中，第1个参数及以后的参数传递给Test03
}

// 不定参数且类型不一样
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	Test() //无参无返回值函数调用

	Test01(10, 20) //函数调用
	Test02(11, 22) //函数调用

	Test03(1)
	Test03(1, 2, 3)

	Test04(4, 5, 6)

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)
}
