package main //包名

//fmt 包实现了格式化 IO（输入/输出）
import (
	"./src/myLib"
	"fmt"
)

func main() {
	//go里面的全局函数要首字母大写
	fmt.Println(myLib2.Add(1, 1))
	fmt.Println(myLib2.Sub(1, 1))
}
