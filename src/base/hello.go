package base

//1) go语言以包作为管理单位
//2) 每个文件必须先声明包
//3) 程序必须有一个main包（重要）

import "fmt"

//fmt 包实现了格式化 IO（输入/输出）

func HelloWord() { //左括号必须和函数名同行

	//打印
	//"hello go"打印到屏幕， Println()会自动换行
	//调用函数，大部分都需要导入包
	/*
		这也是注释， 这是块注释
	*/
	fmt.Println("hello go") //go语言语句结尾是没有分号
	//字符串输出到控制台 最后自动增加换行字符 \n
	fmt.Println("hello itcast")

}
