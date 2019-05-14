package main //包名

import "fmt" //fmt 包实现了格式化 IO（输入/输出）

func main() { //main 函数是每一个可执行程序所必须包含的
	fmt.Println("Hello, World!") //字符串输出到控制台 最后自动增加换行字符 \n

	//在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾

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

	//常量
	const a_con, b_con, c_con = 1, false, "str" //多重赋值
	println(a_con, b_con, c_con)

	//枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	//iota
	/*
	   const (
	       a = iota   //0
	       b          //1
	       c          //2
	       d = "ha"   //独立值，iota += 1
	       e          //"ha"   iota += 1
	       f = 100    //iota +=1
	       g          //100  iota +=1
	       h = iota   //7,恢复计数
	       i          //8
	   )
	   fmt.Println(a,b,c,d,e,f,g,h,i)
	*/
}
