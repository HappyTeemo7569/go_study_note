package main //包名

//fmt 包实现了格式化 IO（输入/输出）
import (
	"./myLib"
	"fmt"
)

func main() { //main 函数是每一个可执行程序所必须包含的
	fmt.Println("Hello, World!") //字符串输出到控制台 最后自动增加换行字符 \n

	//在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾

	//go里面的函数要首字母大写
	myLib.DataType()

	fmt.Println(myLib.Add(1, 1))
}
