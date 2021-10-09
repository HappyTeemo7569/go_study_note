package main

import "fmt"

func DataType() {

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

	//iota

	/**
	字符类型
	*/
	var ch1, ch2, ch3 byte
	ch1 = 'a'  //字符赋值
	ch2 = 97   //字符的ascii码赋值
	ch3 = '\n' //转义字符
	fmt.Printf("ch1 = %c, ch2 = %c, %c", ch1, ch2, ch3)

	//大写转小写，小写转大写, 大小写相差32，小写大
	fmt.Printf("大写：%d， 小写：%d\n", 'A', 'a') //大写：65， 小写：97
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)

	/**
	string
	1、双引号
	2、字符串有1个或多个字符组成
	3、字符串都是隐藏了一个结束符，'\0'
	*/

	var str string                                    // 声明一个字符串变量
	str = "abc"                                       // 字符串赋值
	ch := str[0]                                      // 取字符串的第一个字符
	fmt.Printf("str = %s, len = %d\n", str, len(str)) //内置的函数len()来取字符串的长度
	fmt.Printf("str[0] = %c, ch = %c\n", str[0], ch)

	//`(反引号)括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	str2 := `hello
    mike \n \r测试
    `
	fmt.Println("str2 = ", str2)

	var comv1 complex64 // 由2个float32构成的复数类型
	comv1 = 3.2 + 12i
	comv2 := 3.2 + 12i        // v2是complex128类型
	comv3 := complex(3.2, 12) // v3结果同v2

	fmt.Println(comv1, comv2, comv3)
	//内置函数real(v1)获得该复数的实部
	//通过imag(v1)获得该复数的虚部
	fmt.Println(real(comv1), imag(comv1))
}
