package main

import "fmt"

func main() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s) //1 1 1

	writeS(s)
	fmt.Println(s) // 2 2 2
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	for _, i := range s {
		i++ //局部变量
	}
}

//写入S
func writeS(s []int) {
	for i := range s {
		s[i] += 1 //直接改的底层数组
	}
}
