package main

import "fmt"

func main() {

	/* 标识符
	   mahesh   kumar   abc   move_name   a_123
	   myname50   _temp   j   a23b9   retVal
	*/

	//Go 语言的字符串可以通过 + 实现：
	fmt.Println("Google" + "Runoob")

	/* 关键字
	   break	default	func	interface	select
	   case	defer	go	map	struct
	   chan	else	goto	package	switch
	   const	fallthrough	if	range	type
	   continue	for	import	return	var
	*/

	//变量的声明必须使用空格隔开
	var age int
	age = 3
	println(age)

	//零值就是变量没有做初始化时系统默认设置的值。 和C的初始化一样

	var a string = "Runoob"
	fmt.Println(a)

	/*
	   var a *int
	   var a []int
	   var a map[string] int
	   var a chan int
	   var a func(string) int
	   var a error // error 是接口
	*/

}
